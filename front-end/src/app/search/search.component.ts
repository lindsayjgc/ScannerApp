import { Component, OnInit, ViewChild } from '@angular/core';
import { catchError, of } from 'rxjs';
import { ActivatedRoute } from '@angular/router';
import { MatPaginator, PageEvent } from '@angular/material/paginator';
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
  currentPage: number = 0;
  resultCount: number = 0;

  displayedColumns: string[] = ['name'];
  dataSource = new MatTableDataSource();

  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
    this.searchPage();
  }

  updatePage(event: PageEvent) {
    this.currentPage = event.pageIndex;
    this.loading = true;
    this.searchPage();
  }

  searchPage() {
    this.searchService.search(this.route.snapshot.paramMap.get('option')!, this.route.snapshot.paramMap.get('query')!, this.currentPage + 1)
      .pipe(catchError(err => {
        console.error(err);
        return of();
      }))
      .subscribe((response) => {
        console.log(response);
        this.resultCount = response.count;
        this.dataSource.data = response.products;
        this.loading = false;
      });
  }

  openProductPage(row: any) {
    console.log(row.code);
  }

}
