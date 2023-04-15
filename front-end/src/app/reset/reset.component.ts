import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { catchError, of, tap } from 'rxjs';

import { VerifyService } from '../services/verify.service';

@Component({
  selector: 'app-reset',
  templateUrl: './reset.component.html',
  styleUrls: ['./reset.component.css']
})
export class ResetComponent {
  constructor(private resetMessage: MatSnackBar, private router: Router, private verifyService: VerifyService) { }

  resetForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
  });

  resetPassword() {
    this.verifyService.resetPassword(this.resetForm.value.email!)
      .pipe(catchError(err => {
        this.resetMessage.open(`Error: ${err.error.message}`, '', {
          duration: 5000,
          panelClass: ['login-message-fail'],
        });
        return of();
      }))
      .subscribe((response) => {

      });
  }
}
