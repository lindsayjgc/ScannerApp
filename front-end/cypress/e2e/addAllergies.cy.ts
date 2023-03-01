


// describe('Adding Allergies', () => {
//   it ('should add allergies from the profile page', () => {
//     // Visit the sign up page
//     cy.visit("http://localhost:4200/register");
  
//     // Fill out the sign up form
//     cy.get('input[formControlName="email"]').type("random1267@gmail.com");
//     cy.contains("First Name").click().type("John");
//     cy.contains("Last Name").click().type("Doe");
//     cy.get('input[formControlName="password"]').type("password123");

//     // Submit the form
//     cy.get('button[type="submit"]').click();

//     // Assert that the user is redirected to the setup page
//     cy.url().should('eq', 'http://localhost:4200/setup')

//       const allergies = ['peanuts', 'shellfish', 'gluten']

//       allergies.forEach((allergy) => {
//         cy.get('.allergen-group input')
//           .type(`${allergy}{enter}`)
//           .should('have.value', '')

//         cy.get('.allergen-group mat-chip-row')
//           .contains(allergy)
//           .should('exist')
//       })

//       cy.get('.allergen-group mat-chip-row')
//         .first()
//         .find('button')
//         .click()

//       cy.get('.allergen-group mat-chip-row')
//         .should('have.length', allergies.length - 1)
//     })

//     it('allows submitting the allergies', () => {
//       cy.get('.allergen-group input')
//         .type('peanuts{enter}')
//         .should('have.value', '')

//       cy.get('.setup-card button[type="submit"]')
//         .click()

//       cy.url().should('eq', 'http://localhost:4200/home')

//   it('allows adding a new allergen', () => {
//     cy.visit('/addallergies')
//     cy.get('input[placeholder="New allergen"]').type('Peanuts{enter}')
//     cy.get('.mat-chip-row').should('have.length', 1)
//     cy.get('.mat-chip-row').should('contain', 'Peanuts')
//   })

// });
// })


// describe('Adding Allergies', () => {
//   it('can add allergies', () => {
//     cy.visit('http://localhost:4200/addallergies')
//     cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
//     cy.get('.login-container').contains('Password').parent().find('input').type('password123')
//     cy.get('.login-container').find('button').contains('Login').click()
//     cy.url().should('include', '/home')

//     cy.get('.toolbarContainer').find('button').click()
//     cy.contains('Log out').click()
//     cy.get('.toolbarContainer').find('button').click()
//     cy.should('not.contain.text', 'Log out')
//     cy.url().should('include', '/login')
//   })
// })

// describe('Allergies', () => {
//   beforeEach(() => {
//     cy.visit('http://localhost:4200/addallergies');
//   });

//   it('Can add and remove allergies', () => {

//     cy.get('[matChipInputFor="chipGrid"] input').debug().type('Peanuts{enter}');

//     cy.get('[matChipInputFor="chipGrid"] input').type('Peanuts{enter}');
//     cy.get('.mat-chip-row').contains('Peanuts').find('button').click();

//     cy.get('[matChipInputFor="chipGrid"] input').type('Lactose{enter}');
//     cy.get('.mat-chip-row').contains('Lactose').find('button').click();

//     cy.get('[matChipInputFor="chipGrid"] input').type('Gluten{enter}');
//     cy.get('.mat-chip-row').contains('Gluten').find('button').click();

//     cy.get('button[type="submit"]').click();

//     // Assert that the expected behavior occurs after submitting the form
//     // For example, you could check that the allergies are saved to a database or displayed on the page
//   });
// });

describe('Allergies', () => {

  // beforeEach(() => {
  //   cy.visit('http://localhost:4200/addallergies');
  // });

  it('Can add and remove allergies', () => {

    cy.visit("http://localhost:4200/register");

    cy.get('input[formControlName="email"]').type('randomperson6@gmail.com');
    cy.contains('First Name').click().type('John');
    cy.contains('Last Name').click().type('Doe');
    cy.get('input[formControlName="password"]').type('password123');

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Assert that the user is redirected to the setup page
    cy.url().should('eq', 'http://localhost:4200/setup');

    // Set up allergies
    const allergies = ['peanuts', 'shellfish', 'gluten'];

    allergies.forEach((allergy) => {
      cy.get('.allergen-group input')
        .type(`${allergy}{enter}`)
        .should('have.value', '');

      cy.get('.allergen-group mat-chip-row')
        .contains(allergy)
        .should('exist');
    });

    cy.get('.allergen-group mat-chip-row')
      .first()
      .find('button')
      .click();

    cy.get('.allergen-group mat-chip-row')
      .should('have.length', allergies.length - 1);

    // Submit the allergies
    cy.get('.setup-card button[type="submit"]')
      .click();

    // Assert that the user is redirected to the home page
    cy.url().should('eq', 'http://localhost:4200/home');

    cy.visit("http://localhost:4200/profile")

  cy.get('button[routerlink="/addallergies"]')
  .click()


  
    cy.get('mat-form-field.allergen-group input').type('Treenuts{enter}');
    // cy.get('.mat-chip-grid').contains('Peanuts').find('button').click();

    cy.get('mat-form-field.allergen-group input').type('Lactose{enter}');
    // cy.get('.mat-chip-row').contains('Lactose').find('button').click();

    cy.get('mat-form-field.allergen-group input').type('Lettuce{enter}');
    // cy.get('.mat-chip-row').contains('Gluten').find('button').click();

    cy.get('button[type="submit"]').click();
    cy.url().should('eq', 'http://localhost:4200/profile');

    cy.get('ul').should('contain', 'lactose');
  cy.get('ul').should('contain', 'treenuts');
  cy.get('ul').should('contain', 'lettuce');
  
  });
});


// describe('Allergies', () => {
//   beforeEach(() => {
//     cy.visit('http://localhost:4200/addallergies');
//   });

//   it('Can add and remove allergies', () => {
//     cy.get('mat-form-field.allergen-group input').type('Peanuts{enter}');
//     // cy.get('.mat-chip-grid').contains('Peanuts').find('button').click();

//     cy.get('mat-form-field.allergen-group input').type('Lactose{enter}');
//     // cy.get('.mat-chip-row').contains('Lactose').find('button').click();

//     cy.get('mat-form-field.allergen-group input').type('Gluten{enter}');
//     // cy.get('.mat-chip-row').contains('Gluten').find('button').click();

//     cy.get('button[type="submit"]').click();

//     // Assert that the expected behavior occurs after submitting the form
//     // For example, you could check that the allergies are saved to a database or displayed on the page
//     // cy.url().should('eq', 'http://localhost:4200/profile')
//     cy.get('ul').contains('li', 'Peanuts').should('exist')
//     cy.get('ul').contains('li', 'Lactose').should('exist')
//     cy.get('ul').contains('li', 'Gluten').should('exist')
//   });
// });
