import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VoteOverviewComponent } from './vote-overview.component';

describe('VoteOverviewComponent', () => {
  let component: VoteOverviewComponent;
  let fixture: ComponentFixture<VoteOverviewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VoteOverviewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VoteOverviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
