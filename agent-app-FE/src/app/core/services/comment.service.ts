import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { CompanyRequest } from '../models/request/company-request.model';
import { CompanyUpdate } from '../models/request/company-update.model';
import {CommentResponseModel} from "../models/response/comment-response.model";

@Injectable({
  providedIn: 'root'
})
export class CommentService {
  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(
    private http: HttpClient
  ) { }

  getByCompanyId(id: number): Observable<any>{
    return this.http.get(`${environment.api_url}search/comment/${id}/company`, { headers: this.headers, responseType: 'json' });
  }

  postComment(comment: any): Observable<any>{
    console.log("USAO OVDE")
    return this.http.post(`${environment.api_url}comment`, comment, { headers: this.headers, responseType: 'json' });
  }

  getUserByEmail(email:string): Observable<any> {
    return this.http.get(`${environment.api_url}users?email=${email}`, { headers: this.headers, responseType: 'json' });
  }

  // addCompanyRequest(company: CompanyRequest): Observable<any> {
  //   return this.http.post(`${environment.api_url}company`, company, { headers: this.headers, responseType: 'json' });
  // }
  //
  // updateCompany(company: CompanyUpdate): Observable<any> {
  //   return this.http.put(`${environment.api_url}company`, company, { headers: this.headers, responseType: 'json' });
  // }
  //
  // getCompanyRequests(): Observable<any> {
  //   return this.http.get(`${environment.api_url}companyRequests`, { responseType: 'json' });
  // }
  //
  // approve(id: number, approveRequest: boolean): Observable<any> {
  //   return this.http.post(`${environment.api_url}company/approve`, {id: id, approve: approveRequest}, { headers: this.headers, responseType: 'json' });
  // }
  //
  // getAllCompanies(): Observable<any> {
  //   return this.http.get(`${environment.api_url}company`, { responseType: 'json' });
  // }
}
