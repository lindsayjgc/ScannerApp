import { Component, OnInit } from '@angular/core';
import { catchError, of, tap } from 'rxjs';
import { Router } from '@angular/router';

import { UsersService } from '../services/users.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  constructor(private usersService: UsersService, private router: Router) { }

  ngOnInit() {
    this.usersService.loggedIn()
      .pipe(catchError(err => {
        this.router.navigate(['/login']);
        return of();
      }))
      .subscribe();
  }
}
