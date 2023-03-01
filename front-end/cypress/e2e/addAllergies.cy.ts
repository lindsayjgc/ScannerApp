describe('Add Allergies', () => {

  it('Can add and remove allergies', () => {

    cy.visit("http://localhost:4200/register");

    cy.get('input[formControlName="email"]').type('randomperson1@gmail.com');
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

    cy.visit("http://localhost:4200/profile")

    cy.get('button[routerlink="/addallergies"]')
      .click()

    cy.get('mat-form-field.allergen-group input').type('Treenuts{enter}');
    // cy.get('.mat-chip-grid').contains('Peanuts').find('button').click();

    cy.get('mat-form-field.allergen-group input').type('Lactose{enter}');
    // cy.get('.mat-chip-row').contains('Lactose').find('button').click();

    cy.get('mat-form-field.allergen-group input').type('Lettuce{enter}');
    // cy.get('.mat-chip-row').contains('Gluten').find('button').click();

    cy.get('button[type="submit"]').click();
    cy.url().should('eq', 'http://localhost:4200/profile');

    cy.get('ul').should('contain', 'lactose');
    cy.get('ul').should('contain', 'treenuts');
    cy.get('ul').should('contain', 'lettuce');

  });
});