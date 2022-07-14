import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { JobOfferRequest } from '../models/request/job-offer-request.model';

@Injectable({
  providedIn: 'root'
})
export class JobOfferService {
  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(
    private http: HttpClient
  ) { }

  addJobOffer(jobOffer: JobOfferRequest): Observable<any> {
    return this.http.post(`${environment.api_url}jobOffers`, jobOffer, { headers: this.headers, responseType: 'json' });
  }

  getAllJobOffers(): Observable<any> {
    return this.http.get(`${environment.api_url}jobOffers`, { responseType: 'json' });
  }

  getJobOffersByCompany(id: number): Observable<any> {
    return this.http.get(`${environment.api_url}jobOffers/company/${id}`, { responseType: 'json' });
  }

  deleteJobOffer(id: number): Observable<any> {
    return this.http.delete(`${environment.api_url}jobOffers/${id}`, { responseType: 'json' });
  }

}
