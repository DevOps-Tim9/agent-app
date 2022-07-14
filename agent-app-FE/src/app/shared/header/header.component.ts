import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/core/services/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  permissions = [];

  constructor(
    private router: Router,
    public auth: AuthService,
  ) { }

  ngOnInit(): void {
    this.permissions = this.auth.role;
  }

  logout(): void {
    this.auth.logout();
  }

  goTo(route: string) {
    this.router.navigate([route]);
  }

}
