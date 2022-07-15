import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {CompanyRequestResponse} from '../core/models/response/company-response.model';
import {CompanyService} from '../core/services/company.service';
import {CommentService} from "../core/services/comment.service";
import {CommentResponseModel} from "../core/models/response/comment-response.model";
import {FormControl, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-company-requests',
  templateUrl: './comments.component.html',
  styleUrls: ['./comments.component.scss']
})
export class CommentsComponent implements OnInit {
  company: CompanyRequestResponse;
  comments: CommentResponseModel[];
  user: '';
  stars: number[] = [1, 2, 3, 4, 5];
  dialog: boolean = false;
  starForComment: number[] = [5, 4, 3, 2, 1];
  selectedValue = 3
  postForm = new FormGroup({
    position: new FormControl('', Validators.required),
    interviewProcess: new FormControl('', Validators.required),
    impressions: new FormControl('', Validators.required)
  });


  constructor(
    private route: ActivatedRoute,
    private companyService: CompanyService,
    private commentService: CommentService,
  ) {
  }

  ngOnInit(): void {
    this.selectedValue = 3;
    this.getCompanyRequests();
  }



  getCompanyRequests(): void {
    let id = this.route.snapshot.paramMap.get('id')
    this.companyService.getAllCompanies().subscribe((data: CompanyRequestResponse[]) => {
      this.company = data.filter(comp => {
        return comp['id'] = id;
      })[0]

    });
    this.commentService.getByCompanyId(parseInt(id)).subscribe((data: CommentResponseModel[]) => {
      this.comments = data;
    });
    this.commentService.getUserByEmail(localStorage.mail).subscribe(
      data => {
        this.user = data;
      }
    )

  }

  countStar(star) {
    this.selectedValue = star;
  }


  postTheComment() {
    if(!this.postForm.valid){
      alert('Not valid form');
    }
    let newComment = {
      UserOwnerID: this.user['ID'],
      Position: this.postForm.controls['position'].value,
      InterviewProcess: this.postForm.controls['interviewProcess'].value,
      Description: this.postForm.controls['impressions'].value,
      CompanyID: this.company.ID,
      Rating: this.selectedValue.toString()
    }
    this.commentService.postComment(newComment).subscribe(
      date => {
        this.commentService.getByCompanyId(this.company.ID).subscribe((data: CommentResponseModel[]) => {
          this.comments = data;
        });
        this.postForm.reset()
        this.selectedValue = 3
        alert('Comment are added')

      }
    )
  }


  openDialog() {
    this.dialog = true;

  }
}
