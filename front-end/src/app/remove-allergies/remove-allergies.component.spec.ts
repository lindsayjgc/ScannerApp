import { RemoveAllergiesComponent } from "./remove-allergies.component"

describe('remove-allergies.component.spec.ts', () => {
  it('should allow the user to add and remove allergens', () => {
    cy.mount(RemoveAllergiesComponent)
    // Add an allergen
    cy.get('.allergen-group input').type('Peanuts{enter}');
    cy.get('.allergen-group').should('contain', 'Peanuts');

    // Remove an allergen
    cy.get('.allergen-group mat-chip-row button[aria-label="remove Peanuts"]').click();
    cy.get('.allergen-group').should('not.contain', 'Peanuts');
  });
})