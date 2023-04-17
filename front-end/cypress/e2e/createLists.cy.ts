describe('See Products', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/register');
  });

  it('should check if lists can be created', () => {
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

    cy.visit("http://localhost:4200/profile");

    // const listTitles = ['1', '2', '3'];
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("1");
    cy.contains('Add New List').click();
    cy.contains('1');
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("2");
    cy.contains('Add New List').click();
    cy.contains('2');
    cy.contains('Create New List').click();
    cy.contains("Title").click().type("3");
    cy.contains('Add New List').click();
    cy.contains('3');

    cy.visit('http://localhost:4200/profile')
    cy.contains('Delete account').click()
    cy.contains('Yes').click()
    cy.contains("Password").click().type('password123')
    cy.contains('Delete my account').click()
  });
});