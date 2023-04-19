import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProductService } from '../services/product.service';
import { Allergen } from '../services/allergenparams';

import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { GroceryListService } from '../services/grocery-list.service';
import { AddProductDialogComponent } from '../dialogs/add-product-dialog/add-product-dialog.component';
import { DeleteProductDialogComponent } from '../dialogs/delete-product-dialog/delete-product-dialog.component';
import { FavoritesService } from '../services/favorites.service';
import { CheckIfFavorite } from '../services/favorites.service.spec';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css']
})
export class ProductComponent implements OnInit {
  code: string = "o";
  name: string = "";
  image: string = "";
  ingredients: string = "";
  ingredientsList: any[] = [];
  allergens: Allergen[] = [];
  allergyIngredients: string[] = [];
  commaIngredients: string = "";
  allergic: string = "";
  isAllergic: boolean = false;
  loading: boolean = true;
  imageFound: boolean = false;

  titlesToSend: string = "";
  nameToSend: string = "";
  noIngredients: boolean = false;

  itemFavorited: boolean = false;

  invalidProduct: boolean = false;

  constructor(private route: ActivatedRoute, private http: HttpClient, private productService: ProductService, public dialog: MatDialog, private groceryListService: GroceryListService, public dialog2: MatDialog, private favoriteService: FavoritesService) { }

  ngOnInit() {
    this.code = this.route.snapshot.paramMap.get('code') ?? '';
    console.log(this.code);

    this.favoriteService.checkFavorite(this.code).subscribe((response: any) => {
      console.log(response);
      if (response.isFavorite == "false") {
        this.itemFavorited = false;
      }
      else {
        this.itemFavorited = true;
      }
      console.log(this.itemFavorited);
    });

    const fetchIngredients = async () => {
      this.http.get<any>('https://world.openfoodfacts.org/api/v0/product/' + this.code + '.json')
        .subscribe(response => {
          console.log(response);
          if (response.status == 0) {
            this.invalidProduct = true;
          }
          else {
            this.name = response.product.product_name;
          if (!this.name) {
            this.name = "Unnamed product"
          }
          this.image = response.product.image_front_url;
          if (this.image) {
            this.imageFound = true;
          }
          this.ingredients = response.product.ingredients_text
          if (this.ingredients) {
            this.ingredients = this.ingredients.toLowerCase();
            console.log(this.ingredients);
            this.ingredientsList = this.ingredients.split(", ");
            this.commaIngredients = this.ingredientsList.join(',');
            this.productService.checkForAllergies(this.commaIngredients).subscribe((response: any) => {
              this.allergic = response.allergiesPresent;
              console.log(response);
              if (this.allergic.includes("false")) {
                console.log('You are not allergic to any of the ingredients in this product');
                this.isAllergic = false;
              } else {
                this.allergyIngredients = response.allergies.split(',');
                this.isAllergic = true;
              }
            }, error => {
              console.error(error);
              // handle errors
              this.loading = false;
            });
          } else {
            this.ingredientsList = ["Unknown"];
            this.noIngredients = true;
          }
          }

          this.loading = false;
          return this.ingredientsList;
        });
    }

    fetchIngredients().then((ingredientsList) => {
      console.log(this.ingredientsList);
    });
  }

  addItem(name: string) {
        const dialogRef = this.dialog.open(AddProductDialogComponent);

        dialogRef.afterClosed().subscribe((titles: string []) => {
          if (titles) {
              this.titlesToSend = titles.toString();
              this.nameToSend = name.toLowerCase();
              console.log(this.titlesToSend);
              this.groceryListService.addItemsToList(this.titlesToSend, this.nameToSend).subscribe((response) => {
                console.log(response);
              });
          }
        });
  }

  deleteItem(name: string) {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = name;

    const dialogRef = this.dialog2.open(DeleteProductDialogComponent, dialogConfig);


    dialogRef.afterClosed().subscribe((titles: string []) => {
      if (titles) {
          this.titlesToSend = titles.toString();
          this.nameToSend = name.toLowerCase();
          this.groceryListService.deleteItemsInList(this.titlesToSend, this.nameToSend).subscribe((response) => {
            console.log(response);
          });
      }
    });
  }

  favoriteItem() {
    this.favoriteService.addFavorite(this.name, this.code, this.image).subscribe((response) => {
      console.log(response);
      this.itemFavorited = true;
    });
  }

  unfavoriteItem() {
    this.favoriteService.deleteFavorite(this.code).subscribe((response) => {
      console.log(response);
      this.itemFavorited = false;
    });
  }
}
