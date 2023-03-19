import { Component } from '@angular/core';
import { AllergensService } from '../services/allergens.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-allergies',
  templateUrl: './setup.component.html',
  styleUrls: ['./setup.component.css']
})
export class SetupComponent {
  allergens: Allergen[] = [];

  separatorKeysCodes: number[] = [13, 188]; // Enter and comma keys
  addOnBlur = true;

  constructor(private AllergensService: AllergensService, private router: Router) {}

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

  submitAllergies() {
  const allergyString = this.allergens.map(allergen => allergen.name).join(', ');
  this.AllergensService.addAllergy(allergyString).subscribe((response: any) => {
    console.log(response);
    this.router.navigate(['/home']);
  });
}

}

interface Allergen {
  name: string;
}
