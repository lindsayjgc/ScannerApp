import { Component } from '@angular/core';
import { listParam } from '../services/deleteListparam';
import { GroceryItems, GroceryList } from '../services/grocery-list.service.spec';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { GroceryListService } from '../services/grocery-list.service';
import { CreateListDialogComponent } from '../dialogs/create-list-dialog/create-list-dialog.component';
import { AddItemDialogComponent } from '../dialogs/add-item-dialog/add-item-dialog.component';
import { catchError, of } from 'rxjs';
import { UsersService } from '../services/users.service';

@Component({
  selector: 'app-grocery-lists',
  templateUrl: './grocery-lists.component.html',
  styleUrls: ['./grocery-lists.component.css']
})
export class GroceryListsComponent {

  name: string = '';

  
  titlesParam!: listParam;
  listTitles: string = "";
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

  newItemsArray: string[] = [];

  constructor(private usersService: UsersService, private router: Router, public dialog2: MatDialog, public dialog: MatDialog, private groceryListService: GroceryListService) { }



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
      this.groceryListService.getListTitles().subscribe((titles: any) => {
        console.log(titles);
        this.titlesParam = titles;
        this.listTitles = this.titlesParam.titles;
        if (this.listTitles != "" && this.listTitles != "NONE") {
          this.listTitlesArray = this.listTitles.split(',');
          this.listTitlesArray.forEach((title) => {
            this.groceryListService.getListContents(title).subscribe(
              (contents: GroceryItems) => {
                this.listContents[title] = contents.items.split(",");
                this.listNoItems[title] = false;
              },
              (error: any) => {
                console.error(error);
                this.listNoItems[title] = true;
              }
            );
          });
          this.listTitlesArray.forEach((title) => {
            this.groceryListService.getListContents(title).subscribe({
              next: (contents: GroceryItems) => {
                this.listContents[title] = contents.items.split(",");
                this.listNoItems[title] = false;
                for (let i = 0; i < this.listContents[title].length; i++) {
                  const item = this.listContents[title][i];
                  this.itemChecked[item] = false;
                  if (this.listContents[title][i] == "NONE") {
                    this.listNoItems[title] = true;
                  }
                  console.log(item);
                }
              },
              error: (error: any) => {
                console.error(error);
                this.listNoItems[title] = true;
              }
            });
          });
          
        }
      });
  }

  toggleDropdown(title: string): void {
    this.dropdownOpen[title] = !this.dropdownOpen[title];
  }
  createNewList(newTitle: string): void {
    const dialogRef = this.dialog2.open(CreateListDialogComponent);
  
    dialogRef.afterClosed().subscribe((newTitle: string) => {
      if (newTitle) {
        this.groceryListService.createEmptyList(newTitle).subscribe(() => {
          this.listTitlesArray.push(newTitle);
          this.listContents[newTitle] = [];
        });
        this.groceryListService.getListTitles().subscribe((titles: any) => {
          this.listTitles = titles;
        });
        window.location.reload();
      }
    });
  }
  deleteList(listTitle: string) {
    this.groceryListService.deleteEntireLists(listTitle).subscribe(() => {
      console.log(Response);
    });
    window.location.reload();
  }

  addItems(title: string, newItem: string) {
    const dialogRef = this.dialog.open(AddItemDialogComponent);
  
    dialogRef.afterClosed().subscribe((newItemTemp: GroceryList[]) => {
      if (newItemTemp) {
        for (let i = 0; i < newItemTemp.length; i++) {
          this.newItemsArray.push(newItemTemp[i].item);
        }
        this.newItem = this.newItemsArray.join(',');
        this.groceryListService.addItemsToList(title, this.newItem).subscribe((response) => {
          console.log(response);
          if (this.listNoItems[title]) {
            this.newItemObject.items = newItem;
            this.listContents[title] = [newItem];
            this.listNoItems[title] = false;
          }
          else {
            this.listNoItems[title] = false;
          }

          this.groceryListService.getListTitles().subscribe((titles: any) => {
            console.log(titles);
            this.titlesParam = titles;
            this.listTitles = this.titlesParam.titles;
          });
          
          console.log(this.listTitles);
          this.groceryListService.getListContents("list").subscribe((response) => {
            console.log(response);
          });
          this.groceryListService.getListContents(title).subscribe((items: any) => {
            console.log(title);
            console.log(items);
            this.listContents = items;
          });
        });
        window.location.reload();
      }
    });
  }

  deleteItems(title: string) {
    const itemString = this.selectedItems.toString();
    console.log(itemString);
    this.groceryListService.deleteItemsInList(title, itemString).subscribe((response: any) => {
      console.log(response);
      window.location.reload();
    });
  }

  onItemChecked(item: string) {
    if (this.itemChecked[item]) {
      this.selectedItems.push(item);
    } else {
      this.selectedItems = this.selectedItems.filter((i) => i !== item);
    }
  }
}
