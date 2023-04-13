describe('See Products', () => {
    beforeEach(() => {
      cy.visit('http://localhost:4200/register');
    });
  
    it('should open dialog box for adding items', () => {
      // Fill out the sign up form
      cy.get('input[formControlName="email"]').type('random6753@gmail.com');
      cy.contains('First Name').click().type('John');
      cy.contains('Last Name').click().type('Doe');
      cy.get('input[formControlName="password"]').type('password123');
  
      // Submit the form
      cy.get('button[type="submit"]').click();
  
      // Assert that the user is redirected to the setup page
      cy.url().should('eq', 'http://localhost:4200/setup');
  
      // Set up allergies
      const allergies = ['caffeine', 'sugar', 'gluten'];
  
      allergies.forEach((allergy) => {
        cy.get('.allergen-group input')
          .type(`${allergy}{enter}`)
          .should('have.value', '');
  
        cy.get('.allergen-group mat-chip-row')
          .contains(allergy)
          .should('exist');
      });
  
      // Submit the allergies
      cy.get('.setup-card button[type="submit"]')
        .click();
  
      // Assert that the user is redirected to the home page
      cy.url().should('eq', 'http://localhost:4200/home');

      cy.visit("http://localhost:4200/profile");

      // const listTitles = ['1', '2', '3'];
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("1");
    cy.contains('Add New List').click();
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("2");
    cy.contains('Add New List').click();
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("3");
    cy.contains('Add New List').click();


    cy.contains('1').click(); // Click on dropdown for list 1
    cy.get('button').contains('Add Item').click();
    cy.contains('Add Items to List').should('be.visible')
    cy.contains("Product").click().type("eggs");
    cy.contains("Add To List").click();

    cy.contains('2').click(); // Click on dropdown for list 1
    cy.get('button').contains('Add Item').click();
    cy.contains('Add Items to List').should('be.visible')
    cy.contains("Product").click().type("chips");
    cy.contains("Add To List").click();

    cy.contains('3').click(); // Click on dropdown for list 1
    cy.get('button').contains('Add Item').click();
    cy.contains('Add Items to List').should('be.visible')
    cy.contains("Product").click().type("coffee");
    cy.contains("Add To List").click();

    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.contains("Password").click().type('password123')
    cy.contains('Delete my account').click()
    });
  });