import { Component } from '@angular/core';
import { FavoritesService } from '../services/favorites.service';
import { UsersService } from '../services/users.service';
import { Router } from '@angular/router';
import { FavoritesInfo } from '../services/favorites.service.spec';
import { CheckIfFavorite } from '../services/favorites.service.spec';
import { catchError, of } from 'rxjs';

@Component({
  selector: 'app-favorite-products',
  templateUrl: './favorite-products.component.html',
  styleUrls: ['./favorite-products.component.css']
})
export class FavoriteProductsComponent {
  name: string = '';

  favoritesParam!: FavoritesInfo;
  favName: string = "";
  favCode: string = "";
  favImage: string = "";
  favNameArray: string[] = [];
  favCodeArray: string[] = [];
  favImageArray: string[] = [];
  productRows: { name: string; code: string; image: string; }[] = [];

  
  // titlesParam!: listParam;
  // listTitles: string = "";
  // // @Input() listTitlesArray: string [] = [];
  // // @Input() listContents: { [title: string]: string[] } = {};
  // // @Input() listNoItems: { [title: string]: boolean } = {};
  // listTitlesArray: string [] = [];
  // listContents: { [title: string]: string[] } = {};
  // listNoItems: { [title: string]: boolean } = {};
  // itemChecked: { [itemName: string]: boolean } = {};
  // newTitle: string = "";
  // selectedItems: string[] = [];
  // newItem: string = "";
  // newItemObject!: GroceryItems;
  // secondTitle: string = "";

  constructor(private usersService: UsersService, private router: Router, private favoriteService: FavoritesService) { }



  ngOnInit() {
    this.usersService.loggedIn()
    .pipe(catchError(err => {
      this.router.navigate(['/login']);
      return of();
    }))
    .subscribe();

  this.usersService.getUserData().subscribe((data: any) => {
    this.name = `${data.firstname} ${data.lastname}`;
  })
      this.favoriteService.getFavorites().subscribe((contents: any) => {
        console.log(contents);
        if (contents.status != 204) {
          for (let i = 0; i < contents.length; i++) {
            this.favNameArray.push(contents[i].favorite);
            this.favCodeArray.push(contents[i].code);
            this.favImageArray.push(contents[i].image);
            this.productRows.push({name: contents[i].favorite, code: contents[i].code, image: contents[i].image});
          }
        }
      });
  }

  navigateToProduct(code: string): void {
    console.log(this.productRows);
    this.router.navigate(['/product', code]);
  }
}
