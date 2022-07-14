import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MyErrorStateMatcher } from '../core/error-matchers/ErrorStateMatcher';
import { CompanyUpdate } from '../core/models/request/company-update.model';
import { CompanyRequestResponse } from '../core/models/response/company-response.model';
import { CompanyService } from '../core/services/company.service';
import { Snackbar } from '../shared/snackbar/snackbar';

@Component({
  selector: 'app-edit-company',
  templateUrl: './edit-company.component.html',
  styleUrls: ['./edit-company.component.scss']
})
export class EditCompanyComponent implements OnInit {
  company: CompanyRequestResponse;
  newCompanyForm: FormGroup;
  submitted = false;
  matcher: MyErrorStateMatcher = new MyErrorStateMatcher();

  constructor(
    private formBuilder: FormBuilder,
    private companyService: CompanyService,
    private snackBar: Snackbar,
    private router: Router,
  ) {
    if (!this.router.getCurrentNavigation().extras.state) {
      this.router.navigate(['/home']);
    }
    this.company = this.router.getCurrentNavigation().extras.state.company;

    this.newCompanyForm = this.formBuilder.group({
      name: [this.company.Name, Validators.required],
      contact: [this.company.Contact, Validators.required],
      description: [this.company.Description, Validators.required],
    });
  }

  ngOnInit(): void {
    
  }

  onSubmit(): void {
    this.submitted = true;
    if (this.newCompanyForm.invalid) {
      return;
    }
    const newCompany: CompanyUpdate = { ID: 0, name: '', contact: '', description: '' };
    newCompany.ID = this.company.ID;
    newCompany.name = this.newCompanyForm.value.name;
    newCompany.contact = this.newCompanyForm.value.contact;
    newCompany.description = this.newCompanyForm.value.description;

    this.companyService.updateCompany(newCompany).subscribe((res: any) => {
      this.router.navigate(['/home']);
    },
    error => {
      console.log(error);
    });
  }

  get f(): { [key: string]: AbstractControl; } { return this.newCompanyForm.controls; }

}
