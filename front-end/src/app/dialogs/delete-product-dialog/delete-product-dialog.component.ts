import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { GroceryListService } from 'src/app/services/grocery-list.service';
import { listParam } from 'src/app/services/deleteListparam';
import { GroceryItems } from 'src/app/services/grocery-list.service.spec';

@Component({
  selector: 'app-delete-product-dialog',
  templateUrl: './delete-product-dialog.component.html',
  styleUrls: ['./delete-product-dialog.component.css']
})
export class DeleteProductDialogComponent {

  titlesParam!: listParam;
  listTitles: string = "";
  listTitlesArray: string [] = [];
  itemChecked: { [itemName: string]: boolean } = {};
  selectedItems: string[] = [];
  tempTitle: string = "";
  listContents: { [title: string]: string[] } = {};
  listNoItems: { [title: string]: boolean } = {};
  titlesToShow: string [] = [];


  constructor(private groceryListService: GroceryListService,
    public dialogRef: MatDialogRef<DeleteProductDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: string
  ) {}

  ngOnInit() {
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
              this.titlesToShow = [];
              for (let title in this.listContents) {
                console.log(title);
                for (let element of this.listContents[title]) {
                  console.log(element);
                  if (element == this.data.toLowerCase()) {
                    console.log(title);
                    this.titlesToShow.push(title);
                  }
                }
              }
            },
            (error: any) => {
              console.error(error);
              this.listNoItems[title] = true;
            }
          );
        });

        
        // this.listTitlesArray.forEach((title) => {
        //   for (let i = 0; i < this.listContents[title].length; i++) {
        //     if (this.listContents[title][i] == this.data) {
        //       this.titlesToShow.push(title);
        //     }
        //   }
        // });
        
      }
    });
}
onCancel(): void {
  this.dialogRef.close();
}

onItemChecked(item: string) {
  if (this.itemChecked[item]) {
    this.selectedItems.push(item);
  } else {
    this.selectedItems = this.selectedItems.filter((i) => i !== item);
  }
}

}