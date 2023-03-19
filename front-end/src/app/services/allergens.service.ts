import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { allergenparams } from './allergenparams';
import { Allergen } from '../add-allergies/add-allergies.component';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AllergensService {

  private apiUrl = 'http://localhost:4200/api/add-allergies';
  allergiesToDelete: any;

  constructor(private http: HttpClient) { }

  addAllergy(allergyString: string) {
    return this.http.put(this.apiUrl, { allergies: allergyString });
  }

  private apiUrl2 = 'http://localhost:4200/api/delete-allergies';

  deleteAllergy(allergyString: string) {
    const allergyData = { allergies: allergyString };
    return this.http.delete<allergenparams>(this.apiUrl2, { body: allergyData });
  }
  getAllergens(): Observable<Allergen[]> {
    return this.http.get<Allergen[]>(this.apiUrl);
  }

}
export interface DialogData {
  allergies: string;
}

