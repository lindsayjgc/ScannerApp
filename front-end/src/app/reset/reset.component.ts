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

  sentEmail: boolean = false;
  email: string = "";

  resetForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
  });

  resetForm2 = new FormGroup({
    code: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });

  resetPassword() {
    this.email = this.resetForm.value.email!;
    this.verifyService.resetEmail(this.email)
      .pipe(catchError(err => {
        this.resetMessage.open(`Error: ${err.error.message}`, '', {
          duration: 5000,
          panelClass: ['login-message-fail'],
        });
        return of();
      }))
      .subscribe((response) => {
        this.sentEmail = true;
      });
  }

  resetPassword2() {
    this.verifyService.checkCode(this.resetForm2.value.code!, this.email)
      .pipe(catchError(err => {
        this.resetMessage.open(`Error: ${err.error.message}`, '', {
          duration: 5000,
          panelClass: ['login-message-fail'],
        });
        return of();
      }))
      .subscribe((response) => {
        this.verifyService.resetPassword(this.email, this.resetForm2.value.password!)
          .pipe(catchError(err => {
            this.resetMessage.open(`Error: ${err.error.message}`, '', {
              duration: 5000,
              panelClass: ['login-message-fail'],
            });
            return of();
          }))
          .subscribe((response) => {
            this.router.navigate(['/login']);
          });
      });
  }
}
