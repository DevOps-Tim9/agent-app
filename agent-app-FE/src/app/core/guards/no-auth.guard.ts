import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { AuthService } from 'src/app/core/services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class NoAuthGuard implements CanActivate {

  constructor(
    private router: Router,
    private auth: AuthService
  ) {}

  canActivate(): boolean {
    if (this.auth.isAuthenticated) {
      this.router.navigate(['/home']);
      return false;
    }
    return true;
  }

}
