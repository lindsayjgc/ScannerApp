import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Message } from './message';
import { LoggedInUser } from './loggedInUser';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  constructor(private http: HttpClient) { }

  loginUser(email: string, password: string) {
    return this.http.post<Message>('http://localhost:4200/api/login', { email, password }, { observe: 'response' });
  }

  signupUser(email: string, firstName: string, lastName: string, password: string) {
    return this.http.post<Message>('http://localhost:4200/api/signup', { email, firstName, lastName, password }, { observe: 'response' });
  }

  loggedIn() {
    return this.http.get<LoggedInUser>('http://localhost:4200/api/logged-in');
  }
}

