import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { allergenparams } from './allergenparams.service';

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

  // deleteAllergy(allergyName: string): Observable<any> {
  //   return this.http.delete<allergenparams>('http://localhost:4200/api/delete-allergies', {allergyName});
  // }
  // deleteAllergy(allergiesToDelete: string): Observable<any> {
  //   // const params = new HttpParams().set('allergyName', allergyName);
  //   // const options = {
  //   //   headers: new HttpHeaders({
  //   //     'Content-Type': 'application/json',
  //   //   }),
  //   //   observe: 'body' as const,
  //   //   params: params,
  //   // };
  //   // return this.http.delete<any>('http://localhost:4200/api/delete-allergies', options);

  //   return this.http.delete('http://your-backend-api-url/delete-allergy', {
  //     headers: { 'Content-Type': 'application/json' },
  //     body: JSON.stringify({ allergies: this.allergiesToDelete })
  //   });
  // }
  
  
}
export interface DialogData {
  allergies: string;
}

