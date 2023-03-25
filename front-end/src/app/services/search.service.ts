import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class SearchService {

  constructor(private http: HttpClient) { }

  searchUrl = 'http://localhost:4200/openfood/search.pl?action=process';

  search(option: string, query: string) {
    const headers = new HttpHeaders();
    query = query.toLowerCase().split(' ').join('_');

    return this.http.get<any>(this.searchUrl + '&tagtype_0=' + option + '&tag_contains_0=contains&tag_0=' + query + '&json=true', { 'headers': headers });
  }
}
