import { HomeComponent } from "./home.component"

describe('HomeComponent', () => {
  it('can type and submit', () => {
    cy.mount(HomeComponent)
    cy.get('button').should('be.disabled')
    cy.contains('Category').click()
    cy.contains('Brand').click()
    cy.get('mat-form-field').contains('Search by').parent().should('contain.text', 'Brand')
    cy.get('mat-form-field').contains('Enter search term').click().type('Test')
    cy.get('button').should('be.enabled')
  })

  it('can switch to barcode', () => {
    cy.mount(HomeComponent)
    cy.get('button').should('be.disabled')
    cy.contains('Category').click()
    cy.contains('Barcode').click()
    cy.get('mat-form-field').contains('Search by').parent().should('contain.text', 'Barcode')
    cy.get('mat-form-field').contains('Enter barcode number').click().type('012345')
    cy.get('button').should('be.enabled')
  })
})