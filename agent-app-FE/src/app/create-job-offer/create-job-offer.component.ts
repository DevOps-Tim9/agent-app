import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MyErrorStateMatcher } from '../core/error-matchers/ErrorStateMatcher';
import { JobOfferRequest } from '../core/models/request/job-offer-request.model';
import { CompanyRequestResponse } from '../core/models/response/company-response.model';
import { CompanyService } from '../core/services/company.service';
import { JobOfferService } from '../core/services/job-offer.service';
import { Snackbar } from '../shared/snackbar/snackbar';

@Component({
  selector: 'app-create-job-offer',
  templateUrl: './create-job-offer.component.html',
  styleUrls: ['./create-job-offer.component.scss']
})
export class CreateJobOfferComponent implements OnInit {
  companies: CompanyRequestResponse[];
  newJobOfferForm: FormGroup;
  submitted = false;
  matcher: MyErrorStateMatcher = new MyErrorStateMatcher();

  constructor(
    private formBuilder: FormBuilder,
    private jobOfferService: JobOfferService,
    private companyService: CompanyService,
    private snackBar: Snackbar,
    private router: Router,
  ) {
    this.newJobOfferForm = this.formBuilder.group({
      companyID: [-1, Validators.required],
      position: ['', Validators.required],
      jobDescription: ['', Validators.required],
      dailyActivitiesDescription: ['', Validators.required],
      skills: ['', Validators.required],
      link: ['', Validators.required],
    });
  }

  ngOnInit(): void {
    const auth0id = localStorage.getItem('auth0id');
    this.companyService.getAllCompanies().subscribe((data: CompanyRequestResponse[]) => {
      this.companies = data.filter((c) => '"' + c.OwnerId + '"' === auth0id);
    });
  }

  onSubmit(): void {
    this.submitted = true;
    if (this.newJobOfferForm.invalid) {
      return;
    }
    const newJobOffer: JobOfferRequest = { CompanyID: -1, Position: '', JobDescription: '', DailyActivitiesDescription: '', Skills: '', Link: '' };
    newJobOffer.CompanyID = this.newJobOfferForm.value.companyID;
    newJobOffer.Position = this.newJobOfferForm.value.position;
    newJobOffer.JobDescription = this.newJobOfferForm.value.jobDescription;
    newJobOffer.DailyActivitiesDescription = this.newJobOfferForm.value.dailyActivitiesDescription;
    newJobOffer.Skills = this.newJobOfferForm.value.skills;
    newJobOffer.Link = this.newJobOfferForm.value.link;

    this.jobOfferService.addJobOffer(newJobOffer).subscribe((res: any) => {
      console.log(res);
      this.router.navigate(['/job-offers']);
    },
    error => {
      console.log(error);
    });
  }

  get f(): { [key: string]: AbstractControl; } { return this.newJobOfferForm.controls; }

}
