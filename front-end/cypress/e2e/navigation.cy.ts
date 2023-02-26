describe('navigation when not logged in', () => {
  it('redirects to home', () => {
    cy.visit('http://localhost:4200')
    cy.url().should('include', '/home')
  })

  it('can go to login page', () => {
    cy.visit('http://localhost:4200/home')
    cy.get('.toolbarContainer').find('button').click()
    cy.contains('Log in').click()
    cy.url().should('include', '/login')
  })

  it('redirects to login from profile', () => {
    cy.visit('http://localhost:4200/profile')
    cy.url().should('include', '/login')
  })
})