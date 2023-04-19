import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class SearchService {

  constructor(private http: HttpClient) { }

  searchUrl = 'http://localhost:4200/openfood/search.pl?action=process';

  search(option: string, query: string, page: number) {
    return this.http.get<any>(this.searchUrl + '&tagtype_0=' + option + '&tag_contains_0=contains&tag_0=' + query
      + '&page=' + page + '&json=true');
  }

  recordSearch(query: string) {
    return this.http.post<any>('http://localhost:4200/api/search', { query });
  }
}
