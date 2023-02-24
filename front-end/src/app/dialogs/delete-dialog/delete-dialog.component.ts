import { Component } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { of, catchError } from 'rxjs';

import { UsersService } from 'src/app/services/users.service';

@Component({
  selector: 'app-delete-dialog',
  templateUrl: './delete-dialog.component.html',
  styleUrls: ['./delete-dialog.component.css']
})
export class DeleteDialogComponent {
  constructor(private usersService: UsersService, private dialogRef: MatDialogRef<DeleteDialogComponent>, private errorMessage: MatSnackBar) { }

  onPage1: boolean = true;

  deleteForm = new FormGroup({
    password: new FormControl('', [Validators.required]),
  });


  page2() {
    this.onPage1 = false;
  }

  confirmDeletion() {
    this.usersService.loginUser(this.usersService.loggedInEmail, this.deleteForm.value.password!)
      .pipe(catchError(err => {
        this.errorMessage.open(`Error: ${err.error.message}`, '', {
          duration: 5000,
          panelClass: ['login-message-fail'],
        });
        return of();
      }))
      .subscribe(() => {
        this.dialogRef.close(true);
      });
  }
}
