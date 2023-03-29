describe('delete account', () => {
  it('can go to profile page', () => {
    //Register first
    cy.visit("http://localhost:4200/register");
    cy.get('input[formControlName="email"]').type("random@gmail.com");
    cy.contains("First Name").click().type("John");
    cy.contains("Last Name").click().type("Doe");
    cy.get('input[formControlName="password"]').type("password123");
    cy.get('button[type="submit"]').click();

    //Navigate to profile
    cy.visit('http://localhost:4200/')
    cy.get('.toolbarContainer').find('button').click()
    cy.contains('Profile').click()
    cy.url().should('include', '/profile')
  })

  it('gets error for incorrect password', () => {
    //Log in first
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()

    //Attempt to delete account with incorrect password
    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.contains("Password").click().type('wrongpassword');
    cy.contains('Delete my account').click()

    //Error should display
    cy.get('body').find('mat-snack-bar-container').should('contain.text', 'Incorrect password')
  })

  it('deletes account and gets signed out', () => {
    //Log in first
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()

    //Delete account successfully
    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.contains("Password").click().type('password123')
    cy.contains('Delete my account').click()

    //Should have been logged out
    cy.url().should('include', '/login')
    cy.get('.toolbarContainer').find('button').click()
    cy.should('not.contain.text', 'Log out')
  })
})