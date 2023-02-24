describe('delete account', () => {
  it('can go to profile page', () => {
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()

    cy.get('.toolbarContainer').find('button').click()
    cy.contains('Profile').click()
    cy.url().should('include', '/profile')
  })

  it('gets error for incorrect password', () => {
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()

    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.get('body').find('input').type('wrongpassword')
    cy.contains('Delete my account').click()
    cy.get('body').find('mat-snack-bar-container').should('contain.text', 'Incorrect password')
  })

  it('deletes account and gets signed out', () => {
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()

    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.get('body').find('input').type('password123')
    cy.contains('Delete my account').click()

    cy.url().should('include', '/login')
    cy.get('.toolbarContainer').find('button').click()
    cy.should('not.contain.text', 'Log out')
  })
})