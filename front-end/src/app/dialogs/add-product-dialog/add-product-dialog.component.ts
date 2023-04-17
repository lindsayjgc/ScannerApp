import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { GroceryListService } from 'src/app/services/grocery-list.service';
import { listParam } from 'src/app/services/deleteListparam';

@Component({
  selector: 'app-add-product-dialog',
  templateUrl: './add-product-dialog.component.html',
  styleUrls: ['./add-product-dialog.component.css']
})
export class AddProductDialogComponent {

  titlesParam!: listParam;
  listTitles: string = "";
  listTitlesArray: string [] = [];
  itemChecked: { [itemName: string]: boolean } = {};
  selectedItems: string[] = [];


  constructor(private groceryListService: GroceryListService,
    public dialogRef: MatDialogRef<AddProductDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {}

  ngOnInit() {
      this.groceryListService.getListTitles().subscribe((titles: any) => {
        console.log(titles);
        this.titlesParam = titles;
        this.listTitles = this.titlesParam.titles;
        if (this.listTitles != "" && this.listTitles != "NONE") {
          this.listTitlesArray = this.listTitles.split(',');
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