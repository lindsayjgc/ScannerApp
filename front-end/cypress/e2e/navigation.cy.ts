describe('navigation when not logged in', () => {
  it('redirects to home', () => {
    cy.visit('http://localhost:4200')
    cy.url().should('include', '/home')
  })

  it('redirects to login from profile', () => {
    cy.visit('http://localhost:4200/profile')
    cy.url().should('include', '/login')
  })
})