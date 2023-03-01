import { Component, OnInit } from '@angular/core';
import { catchError, Observable, of, tap } from 'rxjs';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { HttpClient } from '@angular/common/http';

import { UsersService } from '../services/users.service';
import { DeleteDialogComponent } from '../dialogs/delete-dialog/delete-dialog.component';
import { AllergensService } from '../services/allergens.service';
import { DeleteAllergyComponent } from '../dialogs/delete-allergy/delete-allergy.component';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  name: string = '';
  email: string = '';
  password: string = '';
  allergies: string[] = [];
  allergyToRemove: string = '';
  constructor(private usersService: UsersService, private router: Router, public dialog1: MatDialog, private errorMessage: MatSnackBar, private allergensService: AllergensService, public dialog2: MatDialog, public dialog: MatDialog) { }

  ngOnInit() {
    this.usersService.loggedIn()
      .pipe(catchError(err => {
        this.router.navigate(['/login']);
        return of();
      }))
      .subscribe();

    this.usersService.getUserData().subscribe((data) => {
      this.name = `${data.firstname} ${data.lastname}`;
      this.allergies = data.allergies.split(',');
      this.email = `${data.email}`;
      this.password = `${data.password}`;
    })
  }

  openDeleteDialog() {
    const dialogRef = this.dialog1.open(DeleteDialogComponent);

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.usersService.deleteUser()
          .pipe(catchError(err => {
            this.errorMessage.open(`Error: ${err.error.message}`, '', {
              duration: 5000,
              panelClass: ['login-message-fail'],
            });
            return of();
          }))
          .subscribe(() => {
            this.usersService.isLoggedIn = false;
            this.router.navigate(['/login']);
          });
      }
    });
  }


  // getAllergies() {
  //   this.http.get<any>('http://localhost:4200/api/user-info').subscribe((data) => {
  //     this.allergies = data.allergies.split(',');
  //   });
  // }

  // removeSelectedAllergies(selectedAllergies: string[]) {
  //   selectedAllergies.forEach((allergy) => {
  //     this.allergyService.deleteAllergy(allergy).subscribe(() => {
  //       this.getAllergies();
  //     });
  //   });
  // }

  // openDialog() {
  //   const dialogRef = this.dialog.open(DeleteAllergyComponent, {
  //     data: { allergies: this.allergies },
  //   });

  //   dialogRef.afterClosed().subscribe((result: string[]) => {
  //     if (result) {
  //       this.removeSelectedAllergies(result);
  //     }
  //   });
  // }

}
