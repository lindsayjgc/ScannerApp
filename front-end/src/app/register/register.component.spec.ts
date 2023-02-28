import { RegisterComponent } from "./register.component"

//TODO: add a test for signing up the account random@gmail.com with password123
describe('RegisterComponent', () => {
  it('mounts', () => {
    cy.mount(RegisterComponent)
  })
})