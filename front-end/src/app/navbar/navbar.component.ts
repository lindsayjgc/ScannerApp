import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { catchError, of, tap } from 'rxjs';

import { UsersService } from '../services/users.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  constructor(public usersService: UsersService, private router: Router, private errorMessage: MatSnackBar) { }

  ngOnInit() {
    this.usersService.loggedIn()
      .pipe(catchError(err => {
        this.usersService.isLoggedIn = false;
        return of();
      }))
      .subscribe((response) => {
        this.usersService.isLoggedIn = true;
      });
  }

  logout() {
    this.usersService.logoutUser()
      .pipe(catchError(err => {
        this.errorMessage.open(`Error: ${err.error.message}`, '', {
          duration: 5000,
          panelClass: ['login-message-fail'],
        });
        return of();
      }))
      .subscribe((response) => {
        this.usersService.isLoggedIn = false;
        this.router.navigate(['/login']);
      });
  }
}
