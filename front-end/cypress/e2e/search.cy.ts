describe('search functionality', () => {
  it('searches for breakfast cereals', () => {
    //Enter search term
    cy.visit('http://localhost:4200/home')
    cy.get('mat-form-field').contains('Enter search term').click().type('Breakfast Cereals')
    cy.get('button[type="submit"]').click()

    //Verify URL
    cy.url().should('include', '/search/categories/breakfast_cereals/1')

    //Verify results
    cy.get('div[class=mat-mdc-paginator-range-label]', { timeout: 20000 }).should('contain.text', '1 – 24')
    cy.get('table').should('contain.text', 'Image')
      .should('contain.text', 'Product Name')
      .should('contain.text', 'Barcode')

    //Ensure changing pages works
    cy.get('button[aria-label="Next page"]').click()
    cy.url().should('include', '/search/categories/breakfast_cereals/2')
    cy.get('div[class=mat-mdc-paginator-range-label]', { timeout: 20000 }).should('contain.text', '25 – 48')

    cy.get('button[aria-label="Previous page"]').click()
    cy.get('div[class=mat-mdc-paginator-range-label]', { timeout: 20000 }).should('contain.text', '1 – 24')

    cy.visit('http://localhost:4200/search/categories/breakfast_cereals/3')
    cy.get('div[class=mat-mdc-paginator-range-label]', { timeout: 20000 }).should('contain.text', '49 – 72')

    //Return to search bar
    cy.get('button').contains('New search').click()
    cy.url().should('include', '/home')
  })

  it('searches for a specific product', () => {
    //Enter barcode number
    cy.visit('http://localhost:4200/home')
    cy.contains('Category').click()
    cy.contains('Barcode').click()
    cy.get('mat-form-field').contains('Enter barcode number').click().type('8410000001013')
    cy.get('button[type="submit"]').click()

    //Should go directly to product page
    cy.get('div[class=product-details-container]').should('contain.text', 'Chips Ahoy Original')
  })

  it('gets error for nonexistent product', () => {
    //Enter barcode number
    cy.visit('http://localhost:4200/home')
    cy.contains('Category').click()
    cy.contains('Barcode').click()
    cy.get('mat-form-field').contains('Enter barcode number').click().type('5')
    cy.get('button[type="submit"]').click()

    //Get error
    cy.get('div[class=productContainer]').should('contain.text', 'No product found')
  })
})