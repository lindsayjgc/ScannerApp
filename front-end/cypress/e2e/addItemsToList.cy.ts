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
    cy.contains('Create').click();
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("2");
    cy.contains('Create').click();
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("3");
    cy.contains('Create').click();

    cy.contains('1').click(); // Click on dropdown for list 1
    cy.contains("Add Item").click();
    // cy.contains("Item").click({ force: true }).type("eggs");
    // cy.get('input').click().type("eggs");

    cy.contains('Add Items to List').should('be.visible')
    // cy.get('input[ng-model="newItem"]').click().type("eggs");
    // cy.get('Item').click().type("eggs");
    
    // cy.contains("Add").click();
    // cy.contains('2').click(); // Click on dropdown for list 2
    // cy.contains("Add Item").click();
    // cy.contains("Item").click().type("milk");
    // cy.get('input').click();
    // cy.contains('3').click(); // Click on dropdown for list 3
    // cy.contains("Add Item").click();
    // cy.contains("Item").click().type("syrup");
    // cy.contains("Add").click();
    });
  });