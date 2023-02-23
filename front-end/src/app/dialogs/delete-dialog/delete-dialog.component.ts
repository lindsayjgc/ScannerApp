import { Component } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';

import { UsersService } from 'src/app/services/users.service';

@Component({
  selector: 'app-delete-dialog',
  templateUrl: './delete-dialog.component.html',
  styleUrls: ['./delete-dialog.component.css']
})
export class DeleteDialogComponent {
  constructor(private usersService: UsersService, private dialogRef: MatDialogRef<DeleteDialogComponent>) { }

  onPage1: boolean = true;

  deleteForm = new FormGroup({
    password: new FormControl('', [Validators.required]),
  });


  page2() {
    this.onPage1 = false;
  }

  confirmDeletion() {
    this.dialogRef.close();
  }
}
