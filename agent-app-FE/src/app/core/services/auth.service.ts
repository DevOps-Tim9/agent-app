import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import auth0 from 'auth0-js';
import { BehaviorSubject, bindNodeCallback, Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { UserRequest } from '../models/request/user-request.model';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });

  register(user: UserRequest): Observable<any> {
    return this.http.post(`${environment.api_url}register`, user, { headers: this.headers, responseType: 'json' });
  }

  userProfile: any;

  auth0 = new auth0.WebAuth({
    clientID: environment.auth.clientID,
    domain: environment.auth.domain,
    responseType: 'token id_token',
    audience: environment.auth.audience,
    redirectUri: environment.auth.redirect,
    scope: 'openid profile'
  });

  private authFlag = 'isLoggedIn';

  token$ = new BehaviorSubject<string>(null);
  userProfile$ = new BehaviorSubject<any>(null);
  onAuthSuccessUrl = '/home';
  onAuthFailureUrl = '/';
  logoutUrl = environment.auth.logout;

  parseHash$ = bindNodeCallback(this.auth0.parseHash.bind(this.auth0));
  checkSession$ = bindNodeCallback(this.auth0.checkSession.bind(this.auth0));

  constructor(
    public router: Router, 
    private http: HttpClient
  ) {}

  public login(): void {
    this.auth0.authorize();
  }

  handleLoginCallback() {
    if (window.location.hash && !this.isAuthenticated) {
      this.parseHash$().subscribe(
        authResult => {
          this.localLogin(authResult);
          this.router.navigate([this.onAuthSuccessUrl]).then();
        },
        err => this.handleError(err)
      );
    }
  }

  private localLogin(authResult) {
    this.token$.next(authResult.accessToken);
    this.userProfile$.next(authResult.idTokenPayload);
    localStorage.setItem(this.authFlag, JSON.stringify(true));

    const jwt: JwtHelperService = new JwtHelperService();
    localStorage.setItem("role", JSON.stringify(jwt.decodeToken(authResult.accessToken).permissions));
    localStorage.setItem("auth0id", JSON.stringify(jwt.decodeToken(authResult.accessToken).sub));
  }

  get isAuthenticated(): boolean {
    return JSON.parse(localStorage.getItem(this.authFlag));
  }

  get role(): any {
    return JSON.parse(localStorage.getItem("role"));
  }

  renewAuth() {
    if (this.isAuthenticated) {
      this.checkSession$({}).subscribe(
        authResult => {
          this.localLogin(authResult);
        },
        err => {
          localStorage.removeItem(this.authFlag);
          this.router.navigate([this.onAuthFailureUrl]).then();
        }
      );
    }
  }

  private localLogout() {
    localStorage.setItem(this.authFlag, JSON.stringify(false));
    localStorage.setItem("role", JSON.stringify(""));
    this.token$.next(null);
    this.userProfile$.next(null);
  }

  logout() {
    this.localLogout();
    this.auth0.logout({
      returnTo: this.logoutUrl,
      clientID: environment.auth.clientID
    });
  }

  private handleError(err) {
    if (err.error_description) {
      console.error(err.error_description);
    } else {
      console.error(JSON.stringify(err));
    }
  }

  public userHasScopes(scopes: Array<string>): boolean {
    const grantedScopes = JSON.parse(localStorage.getItem('scopes')).split(' ');
    return scopes.every(scope => grantedScopes.includes(scope));
  }
}
