import { ProfileComponent } from "./profile.component"

describe('add-allergies', () => {
  it('should allow the user to add and remove new allergens', () => {
    cy.mount(ProfileComponent)
    // Add an allergen
    cy.get('.allergen-group input').type('Peanuts{enter}', { force: true });
    cy.get('.allergen-group').should('contain', 'Peanuts');

    // Remove an allergen
    cy.get('.allergen-group mat-chip-row button[aria-label="remove Peanuts"]').click();
    cy.get('.allergen-group').should('not.contain', 'Peanuts');
  });
})