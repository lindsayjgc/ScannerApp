import { LoginComponent } from "./login.component"

describe('login.component.spec.ts', () => {
  it('gets error for nonexistent user', () => {
    cy.mount(LoginComponent)
    cy.get('.login-container').contains('Email').parent().find('input').type('fakeemail@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.get('body').find('mat-snack-bar-container').should('contain.text', 'Email not registered to any account')
  })

  it('gets error for wrong password', () => {
    cy.mount(LoginComponent)
    cy.get('.login-container').contains('Email').parent().find('input').type('jd@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.get('body').find('mat-snack-bar-container').should('contain.text', 'Incorrect password')
  })
})