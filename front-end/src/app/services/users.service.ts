import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

import { LoginResponse } from './loginresponse';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  constructor(private http: HttpClient) { }

  loginUser(email: string, password: string) {
    return this.http.post<LoginResponse>('http://localhost:4200/api/login', { email, password });
  }
}
