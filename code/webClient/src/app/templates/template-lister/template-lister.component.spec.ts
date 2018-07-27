import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TemplateListerComponent } from './template-lister.component';

describe('TemplateListerComponent', () => {
  let component: TemplateListerComponent;
  let fixture: ComponentFixture<TemplateListerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TemplateListerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TemplateListerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
