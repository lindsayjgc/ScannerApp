import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { catchError, Observable, of } from 'rxjs';

import { UsersService } from '../services/users.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  constructor(private usersService: UsersService, private loginMessage: MatSnackBar) { }

  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });

  loginUser() {
    this.usersService.loginUser(this.loginForm.value.email!, this.loginForm.value.password!)
      .pipe(catchError((error: any, caught: Observable<any>): Observable<any> => {
        this.loginMessage.open(`Error: ${error.error.message}`, '', {
          duration: 5000,
          panelClass: ['login-message-fail'],
        });
        return of();
      }
      ))
      .subscribe(response => {
        this.loginMessage.open('Login successful!', '', {
          duration: 5000,
          panelClass: ['login-message-success'],
        });
      })
  }
}
