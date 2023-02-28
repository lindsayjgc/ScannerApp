import { NavbarComponent } from "./navbar.component"

describe('NavbarComponent', () => {
  it('can login when logged out', () => {
    cy.mount(NavbarComponent)
    cy.get('.toolbarContainer').find('button').click()
    cy.get('[routerLink="/login"]').should('contain.text', 'Log in')
  })
})