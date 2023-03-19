import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { SetupComponent } from './setup/setup.component';
import { HomeComponent } from './home/home.component';
import { ProfileComponent } from './profile/profile.component';
import { AddAllergiesComponent } from './add-allergies/add-allergies.component';
import { RemoveAllergiesComponent } from './remove-allergies/remove-allergies.component';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'home' },
  { path: 'home', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'setup', component: SetupComponent },
  { path: 'profile', component: ProfileComponent },
  { path: 'addallergies', component: AddAllergiesComponent},
  { path: 'removeallergies', component: RemoveAllergiesComponent}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
