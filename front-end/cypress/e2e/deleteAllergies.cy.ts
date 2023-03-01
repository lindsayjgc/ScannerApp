// describe('Deleting Allergies', () => {
//     it('should remove allergies when "Remove Allergies" button is clicked', () => {
//       // Mock allergies data
//       const allergies = ['Peanuts', 'Shellfish', 'Lactose'];
  
//       // Visit the profile page with allergies
//       cy.visit('http://localhost:4200/profile', {
//         onBeforeLoad(win) {
//           // Stub the Angular component's properties and methods
//           cy.stub(win, '__ngContext__').returns({
//             name: 'John Doe',
//             allergies: allergies,
//             openDialog: cy.stub().as('openDialog'),
//           });
//         },
//       });
  
//       // Find the list of allergies
//       cy.get('ul')
//         .find('li')
//         .should('have.length', allergies.length) // Assert that the list contains the expected number of allergies
//         .each(($li, index) => {
//           // Assert that each allergy in the list matches the expected value
//           cy.wrap($li).should('contain.text', allergies[index]);
//         });
  
//       // Click the "Remove Allergies" button
//       cy.get('button').contains('Remove Allergies').click();
  
//       // Assert that the "openDialog" method was called
//       cy.get('@openDialog').should('have.been.called');
  
//       // Find the list of allergies again
//       cy.get('ul').should('not.exist'); // Assert that the list no longer exists
//     });
//   });
  

// describe('Deleting allergies', () => {
//     const name = 'John';
//     const allergies = ['Peanuts', 'Shellfish'];
  
//     beforeEach(() => {
//       // visit the profile page with allergies

//       it('can go to profile page', () => {
//         cy.visit('http://localhost:4200/login')
//         cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
//         cy.get('.login-container').contains('Password').parent().find('input').type('password123')
//         cy.get('.login-container').find('button').contains('Login').click()
    
//         cy.get('.toolbarContainer').find('button').click()
//         cy.contains('Profile').click()
//         cy.url().should('include', '/profile')
//       })

//       cy.visit('http://localhost:4200/login')
//         cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
//         cy.get('.login-container').contains('Password').parent().find('input').type('password123')
//         cy.get('.login-container').find('button').contains('Login').click()
    
//         cy.get('.toolbarContainer').find('button').click()
//         cy.contains('Profile').click()
//         cy.url().should('include', '/profile')
    
//       cy.visit(`http://localhost:4200/profile`, {
//         onBeforeLoad(win: any) {
//           win.allergies = allergies;
//         },
//       });
//     });
  
//     it('should remove allergies when submitted', () => {
//       // click the "Remove Allergies" button to open the dialog
//       cy.contains('Remove Allergies').click();
  
//       // remove the first allergy from the list
//       const allergyToRemove = allergies[0];
//       cy.get('mat-chip-row')
//         .contains(allergyToRemove)
//         .find('button')
//         .click();
  
//       // submit the form to delete the selected allergy
//       cy.get('form').submit();
  
//       // assert that the allergy was removed from the profile page
//       cy.get('ul').should('not.contain', allergyToRemove);
//     });
  
//     it('should cancel deleting allergies when cancel is clicked', () => {
//       // click the "Remove Allergies" button to open the dialog
//       cy.contains('Remove Allergies').click();
  
//       // click the "Cancel" button to close the dialog
//       cy.contains('Cancel').click();
  
//       // assert that the allergies are still visible on the profile page
//       cy.get('ul').should('contain', allergies[0]).and('contain', allergies[1]);
//     });
//   });
  

describe('Deleting Allergies', () => {
    it('should delete allergies from profile menu', () => {

      // Visit the sign up page
    cy.visit("http://localhost:4200/register");
  
    // Fill out the sign up form
    cy.get('input[formControlName="email"]').type("random12678@gmail.com");
    cy.contains("First Name").click().type("John");
    cy.contains("Last Name").click().type("Doe");
    cy.get('input[formControlName="password"]').type("password123");

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Assert that the user is redirected to the setup page
    cy.url().should("include", "/setup");

    describe('Set up allergies', () => {
      // beforeEach(() => {
      //   cy.visit('/setup')
      // })
    
      it('allows adding and removing allergies', () => {
        const allergies = ['peanuts', 'shellfish', 'gluten']
    
        allergies.forEach((allergy) => {
          cy.get('.allergen-group input')
            .type(`${allergy}{enter}`)
            .should('have.value', '')
    
          cy.get('.allergen-group mat-chip-row')
            .contains(allergy)
            .should('exist')
        })
    
        cy.get('.allergen-group mat-chip-row')
          .first()
          .find('button')
          .click()
    
        cy.get('.allergen-group mat-chip-row')
          .should('have.length', allergies.length - 1)
      })
    
      it('allows submitting the allergies', () => {
        cy.get('.allergen-group input')
          .type('peanuts{enter}')
          .should('have.value', '')
    
        cy.get('.setup-card button[type="submit"]')
          .click()
    
        cy.url().should('include', '/home')
      })
    })   


        cy.visit('/addallergies')
        cy.get('input[placeholder="New allergen"]').type('Peanuts{enter}')
        cy.get('.mat-chip-row').should('have.length', 1)
        cy.get('.mat-chip-row').should('contain', 'Peanuts')
  
      // cy.visit('http://localhost:4200/login')
      // cy.get('.login-container').contains('Email').parent().find('input').type('random@gmail.com')
      // cy.get('.login-container').contains('Password').parent().find('input').type('password123')
      // cy.get('.login-container').find('button').contains('Login').click()
  
      // Submit the form
      cy.get('button[type="submit"]').click();
  
      cy.get('.toolbarContainer').find('button').click()
      cy.contains('Profile').click()
      cy.url().should('include', '/profile')
  
      // Click the Remove Allergies button to open the dialog
      cy.get('button[color="primary"]').click()

      cy.url().should('include', '/removeallergies')
  
        cy.get('input[placeholder="New allergen"]').type('Peanuts{enter}')
        cy.get('.mat-chip-row').should('have.length', 1)
        cy.get('.mat-chip-row').should('contain', 'Peanuts')
  
      // Submit the form to delete the selected allergy
      cy.get('button[type="submit"]').click()
  
  
      // Check that the allergy has been removed from the profile
      cy.get('ul').contains('li', 'Peanuts').should('not.exist')
    })
  })
  

// describe('Deleting Allergies', () => {
//     it('should delete allergies successfully', () => {
      
//       // Navigate to the profile page
//       cy.visit('http://localhost:4200/profile');
  
//       // Click on the 'Remove Allergies' button
//       cy.contains('Remove Allergies').click();
  
//       // Wait for the dialog to appear
//       cy.get('.mat-dialog-container').should('be.visible');
  
//       // Remove the first allergy in the list
//       cy.get('.mat-chip-row').first().contains('button', 'cancel').click();
  
//       // Submit the form to delete the allergies
//       cy.get('form').submit();
  
//       // Wait for the dialog to disappear
//       cy.get('.mat-dialog-container').should('not.exist');
  
//       // Check that the allergies were removed from the profile page
//       cy.get('p').contains('Allergies:');
//       cy.get('li').should('have.length', 1);
//     });
//   });
  