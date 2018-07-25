import { NgModule } from '@angular/core';
import {OptionEditorComponent} from "../../votes/option-editor/option-editor.component";
import {RouterModule, Routes} from "@angular/router";
import {VoteOverviewComponent} from "../../votes/vote-overview/vote-overview.component";
import {VoteViewerComponent} from "../../votes/vote-viewer/vote-viewer.component";

const routes: Routes = [
  { path: '', redirectTo: 'templates', pathMatch: 'full' },
  { path: 'templates', component: VoteOverviewComponent },
  { path: 'templates/:id', component: VoteViewerComponent },
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
