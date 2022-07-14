import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MyErrorStateMatcher } from '../core/error-matchers/ErrorStateMatcher';
import { CompanyRequest } from '../core/models/request/company-request.model';
import { UserRequest } from '../core/models/request/user-request.model';
import { AuthService } from '../core/services/auth.service';
import { CompanyService } from '../core/services/company.service';
import { Snackbar } from '../shared/snackbar/snackbar';

@Component({
  selector: 'app-company-registration',
  templateUrl: './company-registration.component.html',
  styleUrls: ['./company-registration.component.scss']
})
export class CompanyRegistrationComponent implements OnInit {
  newCompanyForm: FormGroup;
  submitted = false;
  matcher: MyErrorStateMatcher = new MyErrorStateMatcher();

  constructor(
    private formBuilder: FormBuilder,
    private companyService: CompanyService,
    private snackBar: Snackbar,
    private router: Router,
  ) {
    this.newCompanyForm = this.formBuilder.group({
      name: ['', Validators.required],
      contact: ['', Validators.required],
      description: ['', Validators.required],
    });
  }

  ngOnInit(): void {

  }

  onSubmit(): void {
    this.submitted = true;
    if (this.newCompanyForm.invalid) {
      return;
    }
    const newCompany: CompanyRequest = { name: '', contact: '', description: '' };
    newCompany.name = this.newCompanyForm.value.name;
    newCompany.contact = this.newCompanyForm.value.contact;
    newCompany.description = this.newCompanyForm.value.description;

    this.companyService.addCompanyRequest(newCompany).subscribe((res: any) => {
      console.log(res);
      this.router.navigate(['/home']);
    },
    error => {
      console.log(error);
    });
  }

  get f(): { [key: string]: AbstractControl; } { return this.newCompanyForm.controls; }
}
