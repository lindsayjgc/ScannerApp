import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeleteAllergyComponent } from './delete-allergy.component';

describe('DeleteAllergyComponent', () => {
  let component: DeleteAllergyComponent;
  let fixture: ComponentFixture<DeleteAllergyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DeleteAllergyComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DeleteAllergyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
