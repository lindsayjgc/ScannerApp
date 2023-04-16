import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { SetupComponent } from './setup/setup.component';
import { HomeComponent } from './home/home.component';
import { ProfileComponent } from './profile/profile.component';
import { SearchComponent } from './search/search.component';
import { ErrorpageComponent } from './errorpage/errorpage.component';
import { ProductComponent } from './product/product.component';
import { ResetComponent } from './reset/reset.component';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'home' },
  { path: 'home', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'setup', component: SetupComponent },
  { path: 'profile', component: ProfileComponent },
  { path: 'reset', component: ResetComponent },
  { path: 'search/:option/:query/:page', component: SearchComponent },
  { path: 'product/:code', component: ProductComponent },
  { path: '**', pathMatch: 'full', component: ErrorpageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
