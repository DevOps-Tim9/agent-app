import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../core/services/auth.service';

@Component({
  selector: 'app-callback',
  templateUrl: './callback.component.html',
  styleUrls: ['./callback.component.scss']
})
export class CallbackComponent implements OnInit {

  constructor(
    public authService: AuthService,
    public router: Router
  ) {}

  ngOnInit() {
    if (!this.authService.isAuthenticated) {
      this.authService.login();
      this.authService.handleLoginCallback();
    } else {
      console.log("uslo ovde");
      this.router.navigate(['/home']).then();
    }
  }

}
