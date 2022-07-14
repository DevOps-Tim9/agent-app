import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MatMenuModule } from '@angular/material/menu';
import { Snackbar } from './snackbar/snackbar';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { HeaderComponent } from './header/header.component';
import {MatToolbarModule} from '@angular/material/toolbar';

@NgModule({
  declarations: [
    HeaderComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    MatMenuModule,
    MatIconModule,
    MatDividerModule,
    MatToolbarModule
  ],
  exports: [
    HeaderComponent,
  ],
  providers: [Snackbar]
})
export class SharedModule { }
