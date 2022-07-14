import { Injectable } from '@angular/core';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';

@Injectable()
export class Snackbar {

  constructor(
    private snackBar: MatSnackBar
  ) { }

  config: MatSnackBarConfig = {
    duration: 3000,
    horizontalPosition: 'center',
    verticalPosition: 'top',
  };

  success(msg: string): void {
    this.config.panelClass = ['success', 'notification'];
    this.snackBar.open(msg, '', this.config);
  }

  error(msg: string): void {
    this.config.panelClass = ['error', 'notification'];
    this.snackBar.open(msg, '', this.config);
  }
}
