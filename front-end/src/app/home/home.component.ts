import { Component } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';

interface Criteria {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  constructor(private router: Router) { }

  searchForm = new FormGroup({
    search: new FormControl('', [Validators.required]),
  });

  searchFormCode = new FormGroup({
    search: new FormControl('', [Validators.required]),
  });

  criterion: Criteria[] = [
    { value: 'categories', viewValue: 'Category' },
    { value: 'brands', viewValue: 'Brand' },
    { value: 'product', viewValue: 'Barcode' },
  ];
  selectedCriteria = this.criterion[0].value;

  query: string = "";
  isBarcode: boolean = false;

  criteriaChange() {
    if (this.selectedCriteria == 'product') {
      this.isBarcode = true;
    }
    else {
      this.isBarcode = false;
    }
  }

  search() {
    if (this.selectedCriteria == 'product') {
      this.query = this.searchFormCode.value.search!;
      this.router.navigate(['/product', this.query]);
    }
    else {
      this.query = this.searchForm.value.search!.toLowerCase().split(' ').join('_');
      this.router.navigate(['/search', this.selectedCriteria, this.query, '1']);
    }
  }
}
