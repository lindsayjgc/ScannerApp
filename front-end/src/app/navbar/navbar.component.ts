import { Component, OnInit } from '@angular/core';
import { MatToolbar } from '@angular/material/toolbar';
import { Router } from '@angular/router';
import { catchError, Observable, of, tap } from 'rxjs';

import { UsersService } from '../services/users.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  constructor(private usersService: UsersService, private router: Router) { }

  loggedIn = false;

  ngOnInit() {
    this.usersService.loggedIn()
      .pipe(catchError(err => {
        this.loggedIn = false;
        return of();
      }),
        tap((response) => {
          this.loggedIn = true;
        })
      )
      .subscribe();
  }

  logout() {
    this.router.navigate(['/login']);
  }
}
