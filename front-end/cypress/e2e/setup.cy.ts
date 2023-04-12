describe('Set Up', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/register');
  });

  it('should sign up a new user and set up allergies', () => {
    // Fill out the sign up form
    cy.get('input[formControlName="email"]').type('random1267987@gmail.com');
    cy.contains('First Name').click().type('John');
    cy.contains('Last Name').click().type('Doe');
    cy.get('input[formControlName="password"]').type('password123');

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Assert that the user is redirected to the setup page
    cy.url().should('eq', 'http://localhost:4200/setup');

    // Set up allergies
    const allergies = ['peanuts', 'shellfish', 'gluten'];

    allergies.forEach((allergy) => {
      cy.get('.allergen-group input')
        .type(`${allergy}{enter}`)
        .should('have.value', '');

      cy.get('.allergen-group mat-chip-row')
        .contains(allergy)
        .should('exist');
    });

    cy.get('.allergen-group mat-chip-row')
      .first()
      .find('button')
      .click();

    cy.get('.allergen-group mat-chip-row')
      .should('have.length', allergies.length - 1);

    // Submit the allergies
    cy.get('.setup-card button[type="submit"]')
      .click();

    // Assert that the user is redirected to the home page
    cy.url().should('eq', 'http://localhost:4200/home');

    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.contains("Password").click().type('password123')
    cy.contains('Delete my account').click()
  });
});


