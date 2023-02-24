//TODO: add a test for signing up the account random@gmail.com with password123

describe('log in and log out', () => {
  it('gets error for nonexistent user', () => {
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('fakeemail@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.get('body').find('mat-snack-bar-container').should('contain.text', 'Email not registered to any account')
  })

  it('gets error for wrong password', () => {
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.get('body').find('mat-snack-bar-container').should('contain.text', 'Incorrect password')
  })

  it('can log in and log out', () => {
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.url().should('include', '/home')

    cy.get('.toolbarContainer').find('button').click()
    cy.contains('Log out').click()
    cy.get('.toolbarContainer').find('button').click()
    cy.should('not.contain.text', 'Log out')
    cy.url().should('include', '/login')
  })
})