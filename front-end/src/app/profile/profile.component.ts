import { Component, OnInit , Input } from '@angular/core';
import { catchError, Observable, of, tap } from 'rxjs';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatExpansionModule } from '@angular/material/expansion';

import { UsersService } from '../services/users.service';
import { DeleteDialogComponent } from '../dialogs/delete-dialog/delete-dialog.component';
import { AllergensService } from '../services/allergens.service';
import { Allergen } from '../services/allergenparams';
import { GroceryListService } from '../services/grocery-list.service';
import { GroceryItems } from '../services/grocery-list.service.spec';
import { CreateListDialogComponent } from '../dialogs/create-list-dialog/create-list-dialog.component';
import { listParam } from '../services/deleteListparam';
import { AddItemDialogComponent } from '../dialogs/add-item-dialog/add-item-dialog.component';
import { GroceryListsComponent } from '../grocery-lists/grocery-lists.component';
import { FavoriteProductsComponent } from '../favorite-products/favorite-products.component';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  name: string = '';
  email: string = '';
  password: string = '';
  allergies: string[] = [];

  allergens: Allergen[] = [];
  separatorKeysCodes: number[] = [13, 188]; // Enter and comma keys
  addOnBlur = true;
  selectedAllergies: string[] = [];

  titlesParam!: listParam;
  listTitles: string = "";
  // @Input() listTitlesArray: string [] = [];
  // @Input() listContents: { [title: string]: string[] } = {};
  // @Input() listNoItems: { [title: string]: boolean } = {};
  listTitlesArray: string [] = [];
  listContents: { [title: string]: string[] } = {};
  listNoItems: { [title: string]: boolean } = {};
  itemChecked: { [itemName: string]: boolean } = {};
  dropdownOpen: { [title: string]: boolean } = {};
  newTitle: string = "";
  selectedItems: string[] = [];
  newItem: string = "";
  newItemObject!: GroceryItems;
  secondTitle: string = "";

  constructor(private usersService: UsersService, private router: Router, public dialog1: MatDialog, private errorMessage: MatSnackBar, private allergensService: AllergensService, public dialog2: MatDialog, public dialog: MatDialog, private groceryListService: GroceryListService) { }
  // ngOnChanges(changes: SimpleChanges): void {
  //   if (changes['listTitlesArray'] && changes['listTitlesArray'].currentValue || changes['listContents'] && changes['listContents'].currentValue) {
  //     this.groceryListService.getListTitles().subscribe((titles: any) => {
  //       this.titlesParam = titles;
  //       this.listTitles = this.titlesParam.titles;
  //       if (this.listTitles != "" && this.listTitles != "NONE") {
  //         this.listTitlesArray = this.listTitles.split(',');
  //         this.listTitlesArray.forEach((title) => {
  //           this.groceryListService.getListContents(title).subscribe(
  //             (contents: any) => {
  //               this.listContents[title] = contents;
  //               this.listNoItems[title] = false;
  //             },
  //             (error: any) => {
  //               console.error(error);
  //               this.listNoItems[title] = true;
  //             }
  //           );
  //         });
  //       }
  //     });
  //   }
  // }  

  ngOnInit() {
    this.usersService.loggedIn()
      .pipe(catchError(err => {
        this.router.navigate(['/login']);
        return of();
      }))
      .subscribe();

    this.usersService.getUserData().subscribe((data) => {
      this.name = `${data.firstname} ${data.lastname}`;
      this.allergies = data.allergies.split(',');
      this.email = `${data.email}`;
      this.password = `${data.password}`;
    })
      // this.groceryListService.getListTitles().subscribe((titles: any) => {
      //   console.log(titles);
      //   this.titlesParam = titles;
      //   this.listTitles = this.titlesParam.titles;
      //   if (this.listTitles != "" && this.listTitles != "NONE") {
      //     this.listTitlesArray = this.listTitles.split(',');
      //     this.listTitlesArray.forEach((title) => {
      //       this.groceryListService.getListContents(title).subscribe(
      //         (contents: GroceryItems) => {
      //           this.listContents[title] = contents.items.split(",");
      //           this.listNoItems[title] = false;
      //         },
      //         (error: any) => {
      //           console.error(error);
      //           this.listNoItems[title] = true;
      //         }
      //       );
      //     });
      //     this.listTitlesArray.forEach((title) => {
      //       this.groceryListService.getListContents(title).subscribe({
      //         next: (contents: GroceryItems) => {
      //           this.listContents[title] = contents.items.split(",");
      //           this.listNoItems[title] = false;
      //           for (let i = 0; i < this.listContents[title].length; i++) {
      //             const item = this.listContents[title][i];
      //             this.itemChecked[item] = false;
      //             console.log(item);
      //           }
      //         },
      //         error: (error: any) => {
      //           console.error(error);
      //           this.listNoItems[title] = true;
      //         }
      //       });
      //     });
          
      //   }
      // });
  }

  openDeleteDialog() {
    const dialogRef = this.dialog1.open(DeleteDialogComponent);

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.usersService.deleteUser()
          .pipe(catchError(err => {
            this.errorMessage.open(`Error: ${err.error.message}`, '', {
              duration: 5000,
              panelClass: ['login-message-fail'],
            });
            return of();
          }))
          .subscribe(() => {
            this.usersService.isLoggedIn = false;
            this.router.navigate(['/login']);
          });
      }
    });
  }

  add(event: any): void {
    const input = event.input;
    const value = event.value.trim();

    // Add the new allergen
    if (value) {
      this.allergens.push({ name: value });
    }

    // Clear the input value
    if (input) {
      input.value = '';
    }
  }

  remove(allergen: Allergen): void {
    const index = this.allergens.indexOf(allergen);

    if (index >= 0) {
      this.allergens.splice(index, 1);
    }
  }

  edit(allergen: Allergen, event: any): void {
    const index = this.allergens.indexOf(allergen);

    if (index >= 0) {
      this.allergens[index].name = event.target.value.trim();
    }
  }

  submitAllergies() {
    const allergyString = this.allergens.map(allergen => allergen.name).join(',');
    console.log(allergyString);
    this.allergensService.addAllergy(allergyString).subscribe((response: any) => {
      console.log(response);
      window.location.reload();
    });
  }

  removeAllergies() {
    const allergyString = this.selectedAllergies.toString();
    this.allergensService.deleteAllergy(allergyString).subscribe((response: any) => {
      console.log(response);
      window.location.reload();
    });
  }

  // toggleDropdown(title: string): void {
  //   this.dropdownOpen[title] = !this.dropdownOpen[title];
  // }
  // createNewList(newTitle: string): void {
  //   const dialogRef = this.dialog2.open(CreateListDialogComponent);
  
  //   dialogRef.afterClosed().subscribe((newTitle: string) => {
  //     if (newTitle) {
  //       this.groceryListService.createEmptyList(newTitle).subscribe(() => {
  //         this.listTitlesArray.push(newTitle);
  //         this.listContents[newTitle] = [];
  //       });
  //       this.groceryListService.getListTitles().subscribe((titles: any) => {
  //         this.listTitles = titles;
  //       });
  //       window.location.reload();
  //     }
  //   });
  // }
  // deleteList(listTitle: string) {
  //   this.groceryListService.deleteEntireLists(listTitle).subscribe(() => {
  //     console.log(Response);
  //   });
  //   window.location.reload();
  // }

  // addItems(title: string, newItem: string) {
  //   const dialogRef = this.dialog.open(AddItemDialogComponent);
  
  //   dialogRef.afterClosed().subscribe((newItem: string) => {
  //     if (newItem) {
  //       this.groceryListService.addItemsToList(title, newItem).subscribe((response) => {
  //         console.log(response);
  //         if (this.listNoItems[title]) {
  //           this.newItemObject.items = newItem;
  //           this.listContents[title] = [newItem];
  //           this.listNoItems[title] = false;
  //         }
  //         else {
  //           // this.listContents[title].push(newItem);
  //           this.listNoItems[title] = false;
  //         }

  //         this.groceryListService.getListTitles().subscribe((titles: any) => {
  //           console.log(titles);
  //           this.titlesParam = titles;
  //           this.listTitles = this.titlesParam.titles;
  //         });
          
  //         console.log(this.listTitles);
  //         this.groceryListService.getListContents("list").subscribe((response) => {
  //           console.log(response);
  //         });
  //         this.groceryListService.getListContents(title).subscribe((items: any) => {
  //           console.log(title);
  //           console.log(items);
  //           this.listContents = items;
  //         });
  //       });
  //       window.location.reload();
  //     }
  //   });
  // }

  // deleteItems(title: string) {
  //   const itemString = this.selectedItems.toString();
  //   console.log(itemString);
  //   this.groceryListService.deleteItemsInList(title, itemString).subscribe((response: any) => {
  //     console.log(response);
  //     window.location.reload();
  //   });
  // }

  // onItemChecked(item: string) {
  //   if (this.itemChecked[item]) {
  //     this.selectedItems.push(item);
  //   } else {
  //     this.selectedItems = this.selectedItems.filter((i) => i !== item);
  //   }
  // }

  viewLists() {
    this.router.navigate(['/lists']);
  }
  viewFavorites() {
    this.router.navigate(['/favorites']);
  }

}

