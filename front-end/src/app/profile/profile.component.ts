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
  allergies: string[] = [];
  allergyToRemove: string = '';
  constructor(private usersService: UsersService, private router: Router, public dialog1: MatDialog, private errorMessage: MatSnackBar, private http: HttpClient, private allergensService: AllergensService, public dialog2: MatDialog, private allergyService: AllergensService, public dialog: MatDialog) { }

  ngOnInit() {
    this.usersService.loggedIn()
      .pipe(catchError(err => {
        this.router.navigate(['/login']);
        return of();
      }))
      .subscribe();

      this.getUserData();
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

  getUserData() {
    this.http.get<any>('http://localhost:4200/api/user-info').subscribe((data) => {
      this.name = `${data.firstname} ${data.lastname}`;
      this.allergies = data.allergies.split(',');
    });
  }

  // updateUserProfile(profile: { allergies: string }): Observable<any> {
  //   return this.usersService.updateUserProfile(profile).pipe(
  //     tap(() => {
  //       // Update the allergies in the component's state
  //       this.allergies = profile.allergies.split(',');
  //     }),
  //     catchError((error) => {
  //       this.errorMessage.open(`Error: ${error.message}`, '', {
  //         duration: 5000,
  //         panelClass: ['login-message-fail'],
  //       });
  //       return of();
  //     })
  //   );
  // }

  // remove(allergy: string) {
  //   this.allergensService.deleteAllergy({ allergies: this.allergies.join(',') })
  //     .subscribe();
  // }

  // removeAllergy() {
  //   // this.http.delete(`http://localhost:4200/api/delete-allergies/${this.allergyToRemove}`).subscribe(
  //   //   () => console.log(`Successfully removed ${allergy} allergy`),
  //   //   error => console.error(`Error removing ${allergy} allergy: ${error}`)
  //   // );
  //   if (this.allergyToRemove) {
  //     this.http.delete(`http://localhost:4200/api/delete-allergies/${this.allergyToRemove}`).subscribe(
  //       () => console.log(`Successfully removed ${this.allergyToRemove} allergy`),
  //       error => console.error(`Error removing ${this.allergyToRemove} allergy: ${error}`)
  //     );
  //   }
  // }

  // openDialog(): void {
  //   const dialogRef = this.dialog2.open(DeleteAllergyComponent, {
  //     width: '250px',
  //     data: { allergies: [] },
  //   });

  //   dialogRef.afterClosed().subscribe((result) => {
  //     if (result && result.length) {
  //       this.removeSelectedAllergies(result);
  //     }
  //   });
  // }

  // removeSelectedAllergies(selectedAllergies: string[]): void {
  //   selectedAllergies.forEach((allergy) => {
  //     const url = `http://localhost:4200/api/delete-allergies/${allergy}`;
  //     this.http.delete(url).subscribe();
  //   });

  // }

  // ngOnInit() {
  //   this.getAllergies();
  // }

  // getAllergies() {
  //   this.allergensService.getAllergies().subscribe((allergies: string[]) => {
  //     this.allergies = allergies;
  //   });
  // }

  getAllergies() {
    this.http.get<any>('http://localhost:4200/api/user-info').subscribe((data) => {
      this.allergies = data.allergies.split(',');
    });
  }

  removeSelectedAllergies(selectedAllergies: string[]) {
    selectedAllergies.forEach((allergy) => {
      this.allergyService.deleteAllergy(allergy).subscribe(() => {
        this.getAllergies();
      });
    });
  }

  openDialog() {
    const dialogRef = this.dialog.open(DeleteAllergyComponent, {
      data: { allergies: this.allergies },
    });

    dialogRef.afterClosed().subscribe((result: string[]) => {
      if (result) {
        this.removeSelectedAllergies(result);
      }
    });
  }

  // allergiesToDelete: string = '';

  // deleteAllergies() {
  //   // Prompt the user to enter the allergies to delete
  //   this.allergiesToDelete = prompt("Enter the allergies to delete, separated by commas:");

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

}
