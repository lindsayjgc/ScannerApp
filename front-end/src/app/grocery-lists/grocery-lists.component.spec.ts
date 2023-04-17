import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroceryListsComponent } from './grocery-lists.component';

describe('GroceryListsComponent', () => {
  let component: GroceryListsComponent;
  let fixture: ComponentFixture<GroceryListsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GroceryListsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GroceryListsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
