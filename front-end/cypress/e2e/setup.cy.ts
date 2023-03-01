describe('Setup Allergies', () => {
    it('should setup allergies for a new user', () => {
  
      // Visit the sign up page
      cy.visit("http://localhost:4200/register");
  
      // Fill out the sign up form
      cy.get('input[formControlName="email"]').type("random1267@gmail.com");
      cy.contains("First Name").click().type("John");
      cy.contains("Last Name").click().type("Doe");
      cy.get('input[formControlName="password"]').type("password123");
  
      // Submit the form
      cy.get('button[type="submit"]').click();
  
      // Assert that the user is redirected to the setup page
      cy.url().should("include", "/setup");
  
      describe('Set up allergies', () => {
        // beforeEach(() => {
        //   cy.visit('/setup') // Replace '/allergies' with the actual URL of the page
        // })
      
        it('allows adding and removing allergies', () => {
          const allergies = ['peanuts', 'shellfish', 'gluten']
      
          allergies.forEach((allergy) => {
            cy.get('.allergen-group input')
              .type(`${allergy}{enter}`)
              .should('have.value', '')
      
            cy.get('.allergen-group mat-chip-row')
              .contains(allergy)
              .should('exist')
          })
      
          cy.get('.allergen-group mat-chip-row')
            .first()
            .find('button')
            .click()
      
          cy.get('.allergen-group mat-chip-row')
            .should('have.length', allergies.length - 1)
        })
      
        it('allows submitting the allergies', () => {
          cy.get('.allergen-group input')
            .type('peanuts{enter}')
            .should('have.value', '')
      
          cy.get('.setup-card button[type="submit"]')
            .click()
      
          cy.url().should('include', '/home')
        })
      })      
    })
  })