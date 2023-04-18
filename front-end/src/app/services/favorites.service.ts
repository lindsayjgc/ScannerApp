import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FavoritesInfo } from './favorites.service.spec';
import { CheckIfFavorite } from './favorites.service.spec';

@Injectable({
  providedIn: 'root'
})
export class FavoritesService {

  constructor(private http: HttpClient) { }

  getFavorites() {
    return this.http.get('http://localhost:4200/api/favorite');
  }

  addFavorite(favorite: string, code: string, image: string) {
    return this.http.post('http://localhost:4200/api/favorite', { favorite, code, image });
  }

  deleteFavorite(code: string) {
    return this.http.delete('http://localhost:4200/api/favorite', { body: JSON.stringify({ Code: code }), headers: { 'Content-Type': 'application/json' } });
  }

  checkFavorite(code: string) {
    return this.http.post<CheckIfFavorite>('http://localhost:4200/api/check-favorite', { code });
  }
}
