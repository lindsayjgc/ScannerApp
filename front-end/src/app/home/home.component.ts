import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { catchError, of } from 'rxjs';

import { SearchService } from '../services/search.service';

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
  constructor(private searchService: SearchService) { }

  searchForm = new FormGroup({
    search: new FormControl('', [Validators.required]),
  });

  criterion: Criteria[] = [
    { value: 'categories', viewValue: 'Category' },
  ];
  selectedCriteria = this.criterion[0].value;

  search() {
    this.searchService.search(this.selectedCriteria, this.searchForm.value.search!)
      .pipe(catchError(err => {
        console.error(err);
        return of();
      }))
      .subscribe((response) => {
        console.log(response);
      });
  }
}
