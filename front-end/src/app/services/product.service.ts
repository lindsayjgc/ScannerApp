import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  constructor(private http: HttpClient) { }

  apiUrl = 'http://localhost:4200/api/check-allergies';

  checkForAllergies(ingredients: string) {
    return this.http.post(this.apiUrl, { ingredients });
  }
}
