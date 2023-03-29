import { Component, OnInit, ViewChild } from '@angular/core';
import { catchError, of } from 'rxjs';
import { Router, ActivatedRoute } from '@angular/router';
import { MatPaginator, PageEvent } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';

import { SearchService } from '../services/search.service';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {
  constructor(private searchService: SearchService, private route: ActivatedRoute, private router: Router) { }
  loading: boolean = true;
  currentPage: number = parseInt(this.route.snapshot.paramMap.get('page')!) - 1;
  resultCount: number = 0;

  displayedColumns: string[] = ['image', 'name', 'barcode'];
  dataSource = new MatTableDataSource();

  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
    this.searchPage();
  }

  updatePage(event: PageEvent) {
    this.currentPage = event.pageIndex;
    this.router.navigate(['/search', this.route.snapshot.paramMap.get('option')!, this.route.snapshot.paramMap.get('query')!, this.currentPage + 1]);
    this.loading = true;
    this.searchPage();
  }

  searchPage() {
    this.searchService.search(this.route.snapshot.paramMap.get('option')!, this.route.snapshot.paramMap.get('query')!, this.currentPage + 1)
      .pipe(catchError(err => {
        console.error(err);
        this.router.navigate(['/home']);
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
    this.router.navigate(['/product', row.code]);
  }

}
