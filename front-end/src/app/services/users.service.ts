import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Message } from './message';
import { LoggedInUser } from './loggedInUser';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  constructor(private http: HttpClient) { }

  isLoggedIn: boolean = false;
  loggedInEmail: string = '';

  loginUser(email: string, password: string) {
    return this.http.post<Message>('http://localhost:4200/api/login', { email, password });
  }

  signupUser(email: string, firstName: string, lastName: string, password: string) {
    return this.http.post<Message>('http://localhost:4200/api/signup', { email, firstName, lastName, password });
  }

  loggedIn() {
    return this.http.get<LoggedInUser>('http://localhost:4200/api/logged-in');
  }

  logoutUser() {
    return this.http.post<LoggedInUser>('http://localhost:4200/api/logout', {});
  }

  deleteUser() {
    return this.http.delete<LoggedInUser>('http://localhost:4200/api/delete-user');
  }

  getUserData() {
    return this.http.get<any>('http://localhost:4200/api/user-info');
  }
}

