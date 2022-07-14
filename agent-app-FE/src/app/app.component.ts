import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from './core/services/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  
  constructor(
    private authService: AuthService,
    private router: Router
  ) {
    this.router.events.subscribe((url: any) => {
      if (this.router.url === '/register') {
        this.showNavbar = false;
      } else {
        this.showNavbar = true;
      }
    });
  }

  ngOnInit() {
    this.authService.renewAuth();
  }

  showNavbar = true;
}
