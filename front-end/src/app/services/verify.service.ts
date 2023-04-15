import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Message } from './message';

@Injectable({
  providedIn: 'root'
})
export class VerifyService {
  constructor(private http: HttpClient) { }

  resetPassword(email: string) {
    return this.http.post<Message>('http://localhost:4200/api/verify/reset', { email });
  }
}
