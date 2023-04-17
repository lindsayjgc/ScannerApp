describe('reset password', () => {
  it('creates an account', () => {
    // Fill out the sign up form
    cy.visit("http://localhost:4200/register");
    cy.get('input[formControlName="email"]').type('test@test.com');
    cy.contains('First Name').click().type('John');
    cy.contains('Last Name').click().type('Doe');
    cy.get('input[formControlName="password"]').type('password123');
    cy.get('button[type="submit"]').click();

    // Verify email
    cy.contains('Verification code').click().type('000000');
    cy.get('button[type="submit"]').click();
    cy.url().should('eq', 'http://localhost:4200/setup');
  })

  it('is able to change password', () => {
    cy.visit('http://localhost:4200/login');
    cy.contains('Forgot password?').click();

    //Send verification
    cy.url().should('include', '/reset');
    cy.contains('Email').click().type('test@test.com');
    cy.contains('Send reset email').click();

    //Change password
    cy.contains('Verification code').click().type('000000');
    cy.contains('New password').click().type('password456');
    cy.contains('Change password').click();

    //Login with new password
    cy.url().should('include', '/login');
    cy.contains('Email').click().type('test@test.com')
    cy.contains('Password').click().type('password456')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.url().should('include', '/home');

    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.contains("Password").click().type('password456')
    cy.contains('Delete my account').click()
  })
})