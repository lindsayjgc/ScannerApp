import { AddAllergiesComponent } from "./add-allergies.component"

describe('add-allergies.component.spec.ts', () => {
  it('tests adding allergies', () => {
    cy.mount(AddAllergiesComponent)
      it('should allow the user to add and remove allergens', () => {
        // Add an allergen
        cy.get('.allergen-group input').type('Peanuts{enter}');
        cy.get('.allergen-group mat-chip-row').should('contain', 'Peanuts');
    
        // Remove an allergen
        cy.get('.allergen-group mat-chip-row button[aria-label="remove Peanuts"]').click();
        cy.get('.allergen-group mat-chip-row').should('not.contain', 'Peanuts');
      });
    
      it('should submit the form when the submit button is clicked', () => {
        // Add an allergen
        cy.get('.allergen-group input').type('Gluten{enter}');
        cy.get('.allergen-group mat-chip-row').should('contain', 'Gluten');
    });
    
  })
})