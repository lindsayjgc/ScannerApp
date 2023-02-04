import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  constructor(private http: HttpClient) { }

  loginUser(email: string, password: string) {
    return this.http.post<any>('http://localhost:4200/api/login', { email, password }, { observe: 'response' });
  }
}
