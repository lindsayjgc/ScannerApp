import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { catchError, Observable, of, tap } from 'rxjs';

import { UsersService } from '../services/users.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  constructor(private usersService: UsersService, private loginMessage: MatSnackBar, private router: Router) { }

  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });

  loginUser() {
    this.usersService.loginUser(this.loginForm.value.email!, this.loginForm.value.password!)
      .pipe(catchError(err => {
        this.loginMessage.open(`Error: ${err.error.message}`, '', {
          duration: 5000,
          panelClass: ['login-message-fail'],
        });
        return of();
      }),
        tap((response) => {
          this.router.navigate(['/home'])
            .then(() => {
              window.location.reload();
            });
        })
      )
      .subscribe();
  }
}
