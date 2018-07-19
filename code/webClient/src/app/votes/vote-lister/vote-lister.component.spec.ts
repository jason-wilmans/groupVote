import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VoteListerComponent } from './vote-lister.component';

describe('VoteListerComponent', () => {
  let component: VoteListerComponent;
  let fixture: ComponentFixture<VoteListerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VoteListerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VoteListerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
