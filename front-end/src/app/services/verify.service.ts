import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Message } from './message';

@Injectable({
  providedIn: 'root'
})
export class VerifyService {
  constructor(private http: HttpClient) { }

  resetEmail(email: string) {
    return this.http.post<Message>('http://localhost:4200/api/verify/reset', { email });
  }

  checkCode(code: string, email: string) {
    return this.http.post<any>('http://localhost:4200/api/check-code', { code, email });
  }

  resetPassword(email: string, password: string) {
    return this.http.post<Message>('http://localhost:4200/api/reset-password', { email, password });
  }
}
