// ***********************************************************
// This example support/component.ts is processed and
// loaded automatically before your test files.
//
// This is a great place to put global configuration and
// behavior that modifies Cypress.
//
// You can change the location of this file or turn off
// automatically serving support files with the
// 'supportFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/configuration
// ***********************************************************

// Import commands.js using ES2015 syntax:
import './commands'

// Alternatively you can use CommonJS syntax:
// require('./commands')

import { Type } from '@angular/core';
import { mount, MountConfig } from 'cypress/angular'
import { HttpClientModule } from '@angular/common/http';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatDialogModule } from '@angular/material/dialog';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

// Augment the Cypress namespace to include type definitions for
// your custom command.
// Alternatively, can be defined in cypress/support/component.d.ts
// with a <reference path="./component" /> at the top of your spec.
declare global {
  namespace Cypress {
    interface Chainable {
      mount: typeof customMount
    }
  }
}

const imports = [
  HttpClientModule,
  MatFormFieldModule,
  MatInputModule,
  MatCardModule,
  MatButtonModule,
  MatSnackBarModule,
  MatToolbarModule,
  MatIconModule,
  MatMenuModule,
  MatDialogModule,
  BrowserAnimationsModule,
  FormsModule,
  ReactiveFormsModule
]

function customMount<T>(component: string | Type<T>, config?: MountConfig<T>) {
  if (!config) {
    config = { imports }
  } else {
    config.imports = [...(config?.imports || []), ...imports]
  }
  return mount<T>(component, config)
}

Cypress.Commands.add('mount', customMount)

// Example use:
// cy.mount(MyComponent)