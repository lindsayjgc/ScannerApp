import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { AllergensService, DialogData } from 'src/app/services/allergens.service';

@Component({
  selector: 'app-delete-allergy',
  templateUrl: './delete-allergy.component.html',
  styleUrls: ['./delete-allergy.component.css']
})
export class DeleteAllergyComponent {

  selectedAllergies: string[] = [];
  allergens: Allergen[] = [];
  inputAllergies: string = '';

  separatorKeysCodes: number[] = [13, 188]; // Enter and comma keys
  addOnBlur = true;

  constructor(private AllergensService: AllergensService,
    public dialogRef: MatDialogRef<DeleteAllergyComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { allergies: string[] }
  ) {
    // Copy the initial data into the selectedAllergies array
    this.selectedAllergies = [...data.allergies];
  }

  add(event: any): void {
    const input = event.input;
    const value = event.value.trim();

    // Add the new allergen
    if (value) {
      this.allergens.push({name: value});
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

  deleteAllergies() {
  const allergyString = this.allergens.map(allergen => allergen.name).join(', ');
  this.AllergensService.deleteAllergy(allergyString).subscribe((response: any) => {
    console.log(response);
  });

  this.dialogRef.close();
}

cancel() {
    this.dialogRef.close();
  }

   removeSelected(): void {
    this.dialogRef.close(this.selectedAllergies);
  }
}

interface Allergen {
  name: string;
}
