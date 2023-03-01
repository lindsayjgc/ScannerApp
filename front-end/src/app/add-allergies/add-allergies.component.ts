import { Component } from '@angular/core';
import { AllergensService } from '../services/allergens.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-allergies',
  templateUrl: './add-allergies.component.html',
  styleUrls: ['./add-allergies.component.css']
})
export class AddAllergiesComponent {
  allergens: Allergen[] = [];
  inputText: string = '';

  separatorKeysCodes: number[] = [13, 188]; // Enter and comma keys
  addOnBlur = true;

  constructor(private AllergensService: AllergensService, private router: Router) {}

  ngOnInit(): void {
    // Get stored value from localStorage
    const storedValue = localStorage.getItem('allergens');
    if (storedValue) {
      this.allergens = JSON.parse(storedValue);
    }
  }

  add(event: any): void {
    const input = event.input;
    const value = event.value.trim();

    // Add the new allergen
    if (value) {
      this.allergens.push({name: value});
      this.saveToLocalStorage();
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
      this.saveToLocalStorage();
    }
  }

  edit(allergen: Allergen, event: any): void {
    const index = this.allergens.indexOf(allergen);

    if (index >= 0) {
      this.allergens[index].name = event.target.value.trim();
      this.saveToLocalStorage();
    }
  }

  saveToLocalStorage() {
    // Save the allergens array to localStorage
    localStorage.setItem('allergens', JSON.stringify(this.allergens));
  }

  submitAllergies() {
    const allergyString = this.allergens.map(allergen => allergen.name).join(', ');
    this.AllergensService.addAllergy(allergyString).subscribe((response: any) => {
      console.log(response);
      // Remove the allergens from localStorage after submitting
      localStorage.removeItem('allergens');
      this.router.navigate(['/profile']);
    });
  }
}

export interface Allergen {
  name: string;
}

