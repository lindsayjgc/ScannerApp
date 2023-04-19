describe("Add User", () => {
  it("allows a user to sign up with valid credentials", () => {
    // Visit the sign up page
    cy.visit("http://localhost:4200/register");

    // Fill out the sign up form
    cy.get('input[formControlName="email"]').type("test@test.com");
    cy.contains("First Name").click().type("John");
    cy.contains("Last Name").click().type("Doe");
    cy.get('input[formControlName="password"]').type("password123");

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Gets error for incorrect verification
    cy.contains('Verification code').click().type('0');
    cy.get('button[type="submit"]').click();
    cy.contains("Error: Incorrect code");

    // Verify email correctly
    cy.contains('Verification code').click().type('00000');
    cy.get('button[type="submit"]').click();

    // Assert that the user is redirected to the setup page
    cy.url().should('eq', 'http://localhost:4200/setup');
  });

  it("gets error when signing up with an existing email address", () => {
    // Visit the sign up page
    cy.visit("http://localhost:4200/register");

    // Fill out the sign up form
    cy.get('input[formControlName="email"]').type("test@test.com");
    cy.contains("First Name").click().type("John");
    cy.contains("Last Name").click().type("Doe");
    cy.get('input[formControlName="password"]').type("password123");

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Assert that an error message is displayed
    cy.contains("Error: Email already registered to an account");
  });

  it('gets error for wrong password', () => {
    //Get error for wrong password
    cy.visit('http://localhost:4200/login')
    cy.contains('Email').click().type('test@test.com')
    cy.contains('Password').click().type('password')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.get('body').find('mat-snack-bar-container').should('contain.text', 'Incorrect password')
  })

  it('can log in and log out', () => {
    //Confirm that logging in goes to home page
    cy.visit('http://localhost:4200/login')
    cy.contains('Email').click().type('test@test.com')
    cy.contains('Password').click().type('password123')
    cy.get('.login-container').find('button').contains('Login').click()
    cy.url().should('include', '/home')

    //Get redirected to home when already logged in
    cy.visit('http://localhost:4200/login')
    cy.url().should('include', '/home')
    cy.visit('http://localhost:4200/register')
    cy.url().should('include', '/home')

    //Confirm that logging out goes to log in page
    cy.get('.toolbarContainer').find('button').click()
    cy.contains('Log out').click()
    cy.get('.toolbarContainer').find('button').click()
    cy.should('not.contain.text', 'Log out')
    cy.url().should('include', '/login')
  })

  it('deletes account', () => {
    cy.visit('http://localhost:4200/login')
    cy.contains('Email').click().type('test@test.com')
    cy.contains('Password').click().type('password123')
    cy.get('.login-container').find('button').contains('Login').click()

    cy.get('.toolbarContainer').find('button').click()
    cy.contains('Profile').click();
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.contains("Password").click().type('password123')
    cy.contains('Delete my account').click()
  })
});
