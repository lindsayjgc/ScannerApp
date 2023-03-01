import { RegisterComponent } from "./register.component"

describe('register.component.spec.ts', () => {
  it('signs up user', () => {
    cy.mount(RegisterComponent)
      cy.get('input[formControlName="email"]').type("email@gmail.com");
      cy.contains("First Name").click().type("John");
      cy.contains("Last Name").click().type("Doe");
      cy.get('input[formControlName="password"]').type("password123");
  })
})