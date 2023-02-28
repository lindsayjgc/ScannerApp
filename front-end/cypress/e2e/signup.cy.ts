describe("Sign up", () => {
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
      cy.url().should("include", "/setup");
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
  });
  