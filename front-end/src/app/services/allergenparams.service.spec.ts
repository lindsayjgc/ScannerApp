import { TestBed } from '@angular/core/testing';

import { AllergenparamsService } from './allergenparams.service';

describe('AllergenparamsService', () => {
  let service: AllergenparamsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AllergenparamsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
