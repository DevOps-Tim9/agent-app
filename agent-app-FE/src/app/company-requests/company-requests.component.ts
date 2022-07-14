import { Component, OnInit } from '@angular/core';
import { CompanyRequestResponse } from '../core/models/response/company-response.model';
import { CompanyService } from '../core/services/company.service';
import { Snackbar } from '../shared/snackbar/snackbar';

@Component({
  selector: 'app-company-requests',
  templateUrl: './company-requests.component.html',
  styleUrls: ['./company-requests.component.scss']
})
export class CompanyRequestsComponent implements OnInit {
  companyRequests: CompanyRequestResponse[];
  displayedColumns: string[] = ['name', 'contact', 'owner', 'description', 'approve', 'reject'];

  constructor(
    private companyService: CompanyService,
    private snackBar: Snackbar
  ) { }

  ngOnInit(): void {
    this.getCompanyRequests();
  }

  getCompanyRequests(): void {
    this.companyService.getCompanyRequests().subscribe((data: CompanyRequestResponse[]) => {
      console.log(data);
      this.companyRequests = data;
    });
  }

  approve(id: number, approve: boolean): void {
    this.companyService.approve(id, approve).subscribe((succ: boolean) => {
      this.getCompanyRequests();
      if (approve) {
        this.snackBar.success('You have successfully approved company request.');
      } else {
        this.snackBar.success('You have successfully rejected company request.');
      }
    }, err => {
      this.snackBar.error(err.error);
    });
  }
}
