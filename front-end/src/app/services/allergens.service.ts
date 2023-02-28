import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AllergensService {

  private apiUrl = 'http://localhost:4200/api/add-allergies';

  constructor(private http: HttpClient) { }

  addAllergy(allergyString: string) {
    return this.http.put(this.apiUrl, { allergies: allergyString });
  }
}
