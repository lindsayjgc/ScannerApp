describe('See Products', () => {
    beforeEach(() => {
      cy.visit('http://localhost:4200/register');
    });
  
    it('should check if product view shows allergies', () => {
      // Fill out the sign up form
      cy.get('input[formControlName="email"]').type('random126@gmail.com');
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

      cy.contains('Search by').click();
      cy.contains('Brand').click();
      cy.contains('Enter search term').click().type('Coca Cola');
      cy.get('button[type="submit"]').click();


      cy.wait(4000);
      cy.contains('0049000028904').click();
      cy.url().should('eq', 'http://localhost:4200/product/0049000028904');
      cy.get('.productContainer').should('exist');

      // Check if the message is displayed for allergic ingredients
      cy.get('.product-ingredients').contains('Caution! Ingredients you are allergic to:').should('be.visible');

      cy.visit('http://localhost:4200/profile')
      cy.contains('Delete account').click()
      cy.contains('Yes').click()
      cy.contains("Password").click().type('password123')
      cy.contains('Delete my account').click()
    });
  });