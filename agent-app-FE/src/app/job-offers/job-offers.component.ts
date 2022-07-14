import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { JobOfferRequest } from '../core/models/request/job-offer-request.model';
import { CompanyRequestResponse } from '../core/models/response/company-response.model';
import { CompanyService } from '../core/services/company.service';
import { JobOfferService } from '../core/services/job-offer.service';
import { JobOfferDetailsComponent } from '../job-offer-details/job-offer-details.component';
import { Snackbar } from '../shared/snackbar/snackbar';

@Component({
  selector: 'app-job-offers',
  templateUrl: './job-offers.component.html',
  styleUrls: ['./job-offers.component.scss']
})
export class JobOffersComponent implements OnInit {
  allJobOffers: JobOfferRequest[] = [];
  myJobOffers: JobOfferRequest[] = [];
  displayedColumns: string[] = ['company', 'position', 'jobDescription', 'details', 'delete'];

  constructor(
    private router: Router,
    private jobOfferService: JobOfferService,
    private companyService: CompanyService,
    public dialog: MatDialog,
    private snackBar: Snackbar,
  ) { }

  ngOnInit(): void {
    this.getJobOffers();
  }

  getJobOffers() {
    this.companyService.getAllCompanies().subscribe((data: CompanyRequestResponse[]) => {
      const auth0id = localStorage.getItem('auth0id');
      const companies = data.filter((c) => '"' + c.OwnerId + '"' === auth0id);
      const companyIDS = companies.map((c) => c.ID);

      this.jobOfferService.getAllJobOffers().subscribe((data: JobOfferRequest[]) => {
        this.allJobOffers = data.filter(offer => !companyIDS.includes(offer.CompanyID));
        this.myJobOffers = data.filter(offer => companyIDS.includes(offer.CompanyID));
        this.allJobOffers = this.allJobOffers.map(jo => {
          const company = companies.find(c => c.ID === jo.CompanyID)
          return { companyName: company.Name, ...jo };
        });
        this.myJobOffers = this.myJobOffers.map(jo => {
          const company = companies.find(c => c.ID === jo.CompanyID)
          return { companyName: company.Name, ...jo };
        });
      });
    });
  }

  openDialog(jobOffer) {
    this.dialog.open(JobOfferDetailsComponent, {
      data: {
        jobOffer: jobOffer,
      },
    });
  }

  delete(jobOffer) {
    this.jobOfferService.deleteJobOffer(jobOffer.ID).subscribe(() => {
      this.snackBar.success("Job offer succesfully deleted.");
      this.getJobOffers();
    });
  }

  goTo(route: string) {
    this.router.navigate([route]);
  }

}
