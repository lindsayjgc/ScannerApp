import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { listParam } from './deleteListparam';
import { GroceryItems } from './grocery-list.service.spec';

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
    // return this.http.delete( 'http://localhost:4200/api/delete-lists', { params: { titles } });
      return this.http.delete('http://localhost:4200/api/delete-lists', { body: JSON.stringify({ Titles: titles }), headers: { 'Content-Type': 'application/json' } });
    
  }
  deleteItemsInList(title: string, items: string) {
    // return this.http.delete('http://localhost:4200/api/delete-list-items', { params: { title , items }});
      const payload = {
          title: title,
          items: items
      };
      return this.http.delete('http://localhost:4200/api/delete-list-items', { body: payload });
  
  }
  getListTitles() {
    return this.http.get('http://localhost:4200/api/get-lists');
  }
  getListContents(title: string) {
    return this.http.post<GroceryItems>('http://localhost:4200/api/get-list', { title });
  }
}
