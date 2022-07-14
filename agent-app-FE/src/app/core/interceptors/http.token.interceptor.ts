import { Injectable } from '@angular/core';
import { HttpEvent, HttpInterceptor, HttpHandler, HttpRequest, HttpHeaders, HttpClient } from '@angular/common/http';
import { EMPTY, Observable } from 'rxjs';
import { filter, mergeMap } from 'rxjs/operators';
import { AuthService } from '../services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class HttpTokenInterceptor implements HttpInterceptor {
  constructor(
    private authService: AuthService
  ) { }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (!this.authService.isAuthenticated) {
      return EMPTY;
    } else {
      return this.authService.token$.pipe(
        filter(token => typeof token === 'string'),
        mergeMap(token => {
          const newReq = req.clone({
            setHeaders: {Authorization: `Bearer ${token}`}
          });
          return next.handle(newReq);
        })
      );
    }
  }
}
