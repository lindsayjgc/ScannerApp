import { Component, OnInit, ViewChild } from '@angular/core';
import { catchError, of } from 'rxjs';
import { ActivatedRoute } from '@angular/router';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';

import { SearchService } from '../services/search.service';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {
  constructor(private searchService: SearchService, private route: ActivatedRoute) { }
  loading: boolean = true;

  displayedColumns: string[] = ['name'];
  dataSource = new MatTableDataSource();

  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.searchService.search(this.route.snapshot.paramMap.get('option')!, this.route.snapshot.paramMap.get('query')!)
      .pipe(catchError(err => {
        console.error(err);
        return of();
      }))
      .subscribe((response) => {
        this.dataSource.data = response.products;
        this.dataSource.paginator = this.paginator;
        this.loading = false;
      });
  }

}
