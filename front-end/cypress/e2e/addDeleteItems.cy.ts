describe('Add Delete Items', () => {
    beforeEach(() => {
      cy.visit('http://localhost:4200/register');
    });
  
    it('should open dialog box for adding items', () => {
      // Fill out the sign up form
      cy.get('input[formControlName="email"]').type('test@test.com');
      cy.contains('First Name').click().type('John');
      cy.contains('Last Name').click().type('Doe');
      cy.get('input[formControlName="password"]').type('password123');
      cy.get('button[type="submit"]').click();
  
      // Verify email
      cy.contains('Verification code').click().type('000000');
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
  
      cy.visit('http://localhost:4200/')
      cy.get('.toolbarContainer').find('button').click()
      cy.contains('Grocery Lists').click()
      cy.url().should('include', '/lists')
  
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
      cy.get('button').contains('Add Items').click();
      cy.contains('Add Items To List').should('be.visible')
      const items = ['peanuts', 'shellfish', 'gluten'];
  
      items.forEach((item) => {
        cy.contains('Enter Items').click()
          .type(`${item}{enter}`)
          .should('have.value', '');
  
        cy.get('.item-group mat-chip-row')
          .contains(item)
          .should('exist');
      });
  
      cy.get('.item-group mat-chip-row')
        .first()
        .find('button')
        .click();
  
      cy.get('.item-group mat-chip-row')
        .should('have.length', items.length - 1);
  
      // Submit the allergies
      cy.contains('button', 'Add To List').click();

      cy.contains('1').click(); // Click on dropdown for list 1


      cy.contains('shellfish').click();
      cy.contains('gluten').click();

      cy.get('button').contains('Delete Checked Items').click()
      cy.contains('1').click(); // Click on dropdown for list 1
      cy.contains('NONE');
          
      cy.get('.toolbarContainer').find('button').click()
      cy.contains('Profile').click();
      cy.contains('Delete account').click();
      cy.contains('Yes').click();
      cy.contains("Password").click().type('password123');
      cy.contains('Delete my account').click();
    });
  });