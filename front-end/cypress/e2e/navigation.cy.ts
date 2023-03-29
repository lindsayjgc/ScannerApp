describe('navigation when not logged in', () => {
  //Confirm that base URL redirects to home
  it('redirects to home', () => {
    cy.visit('http://localhost:4200')
    cy.url().should('include', '/home')
  })

  //Confirm that profile is inacessible when logged out
  it('redirects to login from profile', () => {
    cy.visit('http://localhost:4200/profile')
    cy.url().should('include', '/login')
  })

  //Test 404 page with random URL
  it('should go to 404 page', () => {
    cy.visit('http://localhost:4200/random')
    cy.get('body').should('contain.text', '404')
  })
})