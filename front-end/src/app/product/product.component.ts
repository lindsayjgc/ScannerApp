import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProductService } from '../services/product.service';
import { Allergen } from '../services/allergenparams';

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
  allergyIngredients: string []= [];
  commaIngredients: string = "";
  allergic: string = "";
  isAllergic: boolean = false;
  loading: boolean = true;
  imageFound: boolean = false;

  constructor(private route: ActivatedRoute, private http: HttpClient, private productService: ProductService) { }

  ngOnInit() {
    this.code = this.route.snapshot.paramMap.get('code') ?? '';
    console.log(this.code);

    const fetchIngredients = async () => {
      this.http.get<any>('https://world.openfoodfacts.org/api/v0/product/' + this.code + '.json')
    .subscribe(response => {
      console.log(response);
      this.name = response.product.product_name;
      this.image = response.product.image_front_url;
      if(this.image) {
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
      }
      
      this.loading = false;
      return this.ingredientsList;
    });
    }
    
    fetchIngredients().then((ingredientsList) => {
      console.log(this.ingredientsList);
    });    
    }  
}
