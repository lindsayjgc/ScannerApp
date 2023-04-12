describe("Add User", () => {
  it("allows a user to sign up with valid credentials", () => {
    // Visit the sign up page
    cy.visit("http://localhost:4200/register");

    // Fill out the sign up form
    cy.get('input[formControlName="email"]').type("random@gmail.com");
    cy.contains("First Name").click().type("John");
    cy.contains("Last Name").click().type("Doe");
    cy.get('input[formControlName="password"]').type("password123");

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Assert that the user is redirected to the setup page
    cy.url().should('eq', 'http://localhost:4200/setup');
  });

  it("displays an error message when signing up with an existing email address", () => {
    // Visit the sign up page
    cy.visit("http://localhost:4200/register");

    // Fill out the sign up form
    cy.get('input[formControlName="email"]').type("random@gmail.com");
    cy.contains("First Name").click().type("John");
    cy.contains("Last Name").click().type("Doe");
    cy.get('input[formControlName="password"]').type("password123");

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Assert that an error message is displayed
    cy.contains("Error: Email is already registered to an account");
  });

  it('can log in and log out', () => {
    //Confirm that logging in goes to home page
    cy.visit('http://localhost:4200/login')
    cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
    cy.get('.login-container').contains('Password').parent().find('input').type('password123')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.url().should('include', '/home')

    //Confirm that logging out goes to log in page
    cy.get('.toolbarContainer').find('button').click()
    cy.contains('Log out').click()
    cy.get('.toolbarContainer').find('button').click()
    cy.should('not.contain.text', 'Log out')
    cy.url().should('include', '/login')
  })
});
