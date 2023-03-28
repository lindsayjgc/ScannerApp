import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class GroceryListService {

  constructor(private http: HttpClient) { }

  createEmptyList(title: string) {
    return this.http.put('http://localhost:4200/api/create-list', { title });
  }
  addItemsToList(title: string, items: string) {
    return this.http.post('http://localhost:4200/api/add-list-items', { title, items });
  }
  deleteEntireLists(titles: string) {
    // return this.http.delete('http://localhost:4200/api/delete-lists', { titles });
    return this.http.delete('http://localhost:4200/api/delete-lists', { 
    headers: new HttpHeaders().set('Content-Type', 'application/json'), 
    params: { titles }
    });
  }
  deleteItemsInList(title: string, items: string) {
    return this.http.delete('http://localhost:4200/api/delete-list-items', { 
    headers: new HttpHeaders().set('Content-Type', 'application/json'), 
    params: { title , items }
    });
  }
  getListTitles() {
    return this.http.get('http://localhost:4200/api/get-lists');
  }
  getListContents(title: string) {
    return this.http.get('http://localhost:4200/api/get-list', { 
    headers: new HttpHeaders().set('Content-Type', 'application/json'), 
    params: { title }
    });
  }
}
