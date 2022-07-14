import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CallbackComponent } from './callback/callback.component';
import { HomepageComponent } from './homepage/homepage.component';
import { RegisterComponent } from './register/register.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatIconModule } from '@angular/material/icon'
import { MatButtonModule } from '@angular/material/button';
import { ReactiveFormsModule } from '@angular/forms';
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import { CoreModule } from './core/core.module';
import { SharedModule } from './shared/shared.module';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatRadioModule } from '@angular/material/radio';
import {MatSelectModule} from '@angular/material/select';
import { CompanyRegistrationComponent } from './company-registration/company-registration.component';
import { CompanyRequestsComponent } from './company-requests/company-requests.component';
import {MatTableModule} from '@angular/material/table';
import { EditCompanyComponent } from './edit-company/edit-company.component';
import { JobOffersComponent } from './job-offers/job-offers.component';
import { CreateJobOfferComponent } from './create-job-offer/create-job-offer.component';
import {MatTabsModule} from '@angular/material/tabs';
import {MatDialogModule} from '@angular/material/dialog';
import { JobOfferDetailsComponent } from './job-offer-details/job-offer-details.component';

@NgModule({
  declarations: [
    AppComponent,
    CallbackComponent,
    HomepageComponent,
    RegisterComponent,
    CompanyRegistrationComponent,
    CompanyRequestsComponent,
    EditCompanyComponent,
    JobOffersComponent,
    CreateJobOfferComponent,
    JobOfferDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatCardModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatButtonModule,
    MatIconModule,
    MatInputModule,
    MatSnackBarModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatRadioModule,
    MatSelectModule,
    MatTableModule,
    MatTabsModule,
    MatDialogModule,
    CoreModule,
    SharedModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
