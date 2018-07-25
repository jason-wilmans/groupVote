import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VoteViewerComponent } from './vote-viewer.component';

describe('VoteViewerComponent', () => {
  let component: VoteViewerComponent;
  let fixture: ComponentFixture<VoteViewerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VoteViewerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VoteViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
