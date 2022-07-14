import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { skipWhile, take } from 'rxjs/operators';
import { CompanyRequestResponse } from '../core/models/response/company-response.model';
import { AuthService } from '../core/services/auth.service';
import { CompanyService } from '../core/services/company.service';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {
  permissions = [];
  auth0id = '';
  companies: CompanyRequestResponse[];
  displayedColumns: string[] = ['name', 'contact', 'owner', 'description', 'edit'];

  constructor(
    private authService: AuthService,
    private companyService: CompanyService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.permissions = this.authService.role;
    this.auth0id = localStorage.getItem('auth0id');
    this.authService.token$.pipe(
      skipWhile(value => !value),
      take(1))
      .subscribe(value => console.log(value));
    this.companyService.getAllCompanies().subscribe((data: CompanyRequestResponse[]) => {
      if (this.permissions.includes('update:company')) {
        this.companies = data.filter((c) => '"' + c.OwnerId + '"' === this.auth0id);
      } else {
        this.companies = data;
      }
    });
  }

  edit(c: CompanyRequestResponse) {
    this.router.navigate(['/edit-company'], { state: { company: c } });
  }

  goTo(route: string) {
    this.router.navigate([route]);
  }

}
