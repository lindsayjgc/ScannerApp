import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { UsersService } from '../services/users.service';
import { LoginResponse } from '../services/loginresponse';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  constructor(private usersService: UsersService) { }

  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });

  loginUser() {
    this.usersService.loginUser(this.loginForm.value.email!, this.loginForm.value.password!).subscribe((response: LoginResponse) => {
      console.log(response);
    })
  }
}
