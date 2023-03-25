import { Component, OnInit } from '@angular/core';
import { catchError, of } from 'rxjs';
import { ActivatedRoute } from '@angular/router';

import { SearchService } from '../services/search.service';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {
  constructor(private searchService: SearchService, private route: ActivatedRoute) { }

  ngOnInit() {
    // this.searchService.search(this.route.snapshot.paramMap.get('option')!, this.route.snapshot.paramMap.get('query')!)
    //   .pipe(catchError(err => {
    //     console.error(err);
    //     return of();
    //   }))
    //   .subscribe((response) => {
    //     console.log(response);
    //   });
  }

}
