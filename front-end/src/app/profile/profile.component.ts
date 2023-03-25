import { Component, OnInit } from '@angular/core';
import { catchError, Observable, of, tap } from 'rxjs';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';

import { UsersService } from '../services/users.service';
import { DeleteDialogComponent } from '../dialogs/delete-dialog/delete-dialog.component';
import { AllergensService } from '../services/allergens.service';
import { Allergen } from '../services/allergenparams';

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

  allergens: Allergen[] = [];
  separatorKeysCodes: number[] = [13, 188]; // Enter and comma keys
  addOnBlur = true;
  selectedAllergies: string[] = [];

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

  add(event: any): void {
    const input = event.input;
    const value = event.value.trim();

    // Add the new allergen
    if (value) {
      this.allergens.push({ name: value });
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
    this.allergensService.addAllergy(allergyString).subscribe((response: any) => {
      console.log(response);
      window.location.reload();
    });
  }

  removeAllergies() {
    const allergyString = this.selectedAllergies.toString();
    this.allergensService.deleteAllergy(allergyString).subscribe((response: any) => {
      console.log(response);
      window.location.reload();
    });
  }

}

