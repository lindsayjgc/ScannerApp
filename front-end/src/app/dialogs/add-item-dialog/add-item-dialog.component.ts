import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { GroceryListService } from 'src/app/services/grocery-list.service';
import { GroceryList } from 'src/app/services/grocery-list.service.spec';

@Component({
  selector: 'app-add-item-dialog',
  templateUrl: './add-item-dialog.component.html',
  styleUrls: ['./add-item-dialog.component.css']
})
export class AddItemDialogComponent {
  newItem: string = "";
  newItems: GroceryList[] = [];
  separatorKeysCodes: number[] = [13, 188]; // Enter and comma keys
  addOnBlur = true;

  constructor(
    public dialogRef: MatDialogRef<AddItemDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {}

  add(event: any): void {
    const input = event.input;
    const value = event.value.trim();

    // Add the new allergen
    if (value) {
      this.newItems.push({ item: value });
    }

    // Clear the input value
    if (input) {
      input.value = '';
    }
  }

  remove(itemName: GroceryList): void {
    const index = this.newItems.indexOf(itemName);

    if (index >= 0) {
      this.newItems.splice(index, 1);
    }
  }

  edit(itemName: GroceryList, event: any): void {
    const index = this.newItems.indexOf(itemName);

    if (index >= 0) {
      this.newItems[index].item = event.target.value.trim();
    }
  }

  // submitAllergies() {
  //   const allergyString = this.newItems.map(item => item).join(',');
  //   this.AllergensService.addAllergy(allergyString).subscribe((response: any) => {
  //     console.log(response);
  //     this.router.navigate(['/home']);
  //   });
  // }

  onCancel(): void {
    this.dialogRef.close();
  }
}
