import { NgModule } from '@angular/core';
import {OptionEditorComponent} from "../../templates/option-editor/option-editor.component";
import {RouterModule, Routes} from "@angular/router";
import {TemplateOverviewComponent} from "../../templates/template-overview/template-overview.component";
import {TemplateViewerComponent} from "../../templates/template-viewer/template-viewer.component";

const routes: Routes = [
  { path: '', redirectTo: 'templates', pathMatch: 'full' },
  { path: 'templates', component: TemplateOverviewComponent },
  { path: 'templates/:id', component: TemplateViewerComponent },
  { path: 'templates/:id/options', component: OptionEditorComponent }
];

@NgModule({
  imports: [
    RouterModule.forRoot(
      routes,
      { enableTracing: false }
    )
  ],
  declarations: [
  ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
