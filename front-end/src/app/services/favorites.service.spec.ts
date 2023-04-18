export interface FavoritesInfo {
  favorite: string[];
  code: string[];
  image: string[];
}

export interface CheckIfFavorite {
  code: string;
  isFavorite: boolean;
}