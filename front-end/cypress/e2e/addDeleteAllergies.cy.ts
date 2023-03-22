describe('Add Allergies', () => {

  it('Can add and remove allergies', () => {

    cy.visit("http://localhost:4200/register");

    cy.get('input[formControlName="email"]').type('test@gmail.com');
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

    // See allergies on profile page
    cy.visit("http://localhost:4200/profile");

    cy.get('mat-selection-list').should('contain', 'shellfish');
    cy.get('mat-selection-list').should('contain', 'gluten');

    // Add new allergies
    cy.get('mat-form-field.allergen-group input').type('treenuts{enter}');
    cy.get('mat-form-field.allergen-group input').type('lactose{enter}');
    cy.get('mat-form-field.allergen-group input').type('lettuce{enter}');

    cy.get('button[type="submit"]').click();

    cy.get('mat-selection-list').should('contain', 'treenuts');
    cy.get('mat-selection-list').should('contain', 'lactose');
    cy.get('mat-selection-list').should('contain', 'lettuce');

    // Remove allergies
    cy.get('mat-list-option').contains('gluten').click();
    cy.get('mat-list-option').contains('lettuce').click();

    cy.get('button').contains('Remove selected allergies').click()

    cy.get('mat-selection-list').should('not.contain', 'gluten');
    cy.get('mat-selection-list').should('not.contain', 'lettuce');
  });
});