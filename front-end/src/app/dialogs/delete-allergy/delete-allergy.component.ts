import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { AllergensService, DialogData } from 'src/app/services/allergens.service';
import { ProfileComponent } from 'src/app/profile/profile.component';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-delete-allergy',
  templateUrl: './delete-allergy.component.html',
  styleUrls: ['./delete-allergy.component.css']
})
export class DeleteAllergyComponent {
  // selectedAllergies: string[] = [];

  // constructor(
  //   public dialogRef: MatDialogRef<DeleteAllergyComponent>,
  //   @Inject(MAT_DIALOG_DATA) public data: { allergies: string[] }
  // ) {
  //   // Copy the initial data into the selectedAllergies array
  //   this.selectedAllergies = [...data.allergies];
  // }

  // toggleAllergySelection(allergy: string): void {
  //   if (this.selectedAllergies.includes(allergy)) {
  //     this.selectedAllergies = this.selectedAllergies.filter(
  //       (a) => a !== allergy
  //     );
  //   } else {
  //     this.selectedAllergies.push(allergy);
  //   }
  // }

  // cancel(): void {
  //   this.dialogRef.close();
  // }

  // removeSelected(): void {
  //   this.dialogRef.close(this.selectedAllergies);
  // }

  // selectedAllergies: string[] = [];
  // inputAllergies: string = '';
  // allergiesToDelete: any;

  // constructor(private http: HttpClient,
  //   public dialogRef: MatDialogRef<DeleteAllergyComponent>,
  //   @Inject(MAT_DIALOG_DATA) public data: DialogData
  // ) {}

  // toggleAllergySelection(allergy: string) {
  //   if (this.selectedAllergies.includes(allergy)) {
  //     this.selectedAllergies = this.selectedAllergies.filter((a) => a !== allergy);
  //   } else {
  //     this.selectedAllergies.push(allergy);
  //   }
  // }

  // removeAllergy(allergiesToDelete: string) {
  //   // const url = `https://your-backend-api.com/allergies/${allergy}`;
  //   // return this.http.delete(url);

  //   // Send a DELETE request to the backend API
  //   this.http.delete('http://your-backend-api-url/delete-allergy', {
  //     headers: { 'Content-Type': 'application/json' },
  //     body: JSON.stringify({ allergies: this.allergiesToDelete })
  //   }).subscribe(response => {
  //     // Handle the response from the backend API
  //     console.log(response);
  //     alert('Allergies deleted successfully');
  //   }, error => {
  //     // Handle errors
  //     console.log(error);
  //     alert('Error deleting allergies');
  //   });

    
  // }
  

  // removeSelected() {
  //   let allergiesToRemove = this.selectedAllergies;
  
  //   if (this.inputAllergies && this.inputAllergies.trim() !== '') {
  //     const inputAllergiesArray = this.inputAllergies.split(',').map((s) => s.trim());
  //     allergiesToRemove = [...allergiesToRemove, ...inputAllergiesArray];
  //   }
  
  //   allergiesToRemove.forEach((allergy) => {
  //     this.removeAllergy(allergy)
  //   });
  
  //   this.dialogRef.close(allergiesToRemove);
  // }
  

  // cancel() {
  //   this.dialogRef.close();
  // }

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
