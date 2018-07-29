import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from "./infrastructure/app-routing/app-routing-module";

import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatCardModule} from '@angular/material/card';
import {FormsModule} from '@angular/forms';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatListModule} from '@angular/material/list';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';

import { AppComponent } from './app.component';
import { TemplateEditorComponent } from './templates/template-editor/template-editor.component';
import { TemplateListerComponent } from './templates/template-lister/template-lister.component';
import { OptionEditorComponent } from './templates/option-editor/option-editor.component';
import { TemplateOverviewComponent } from './templates/template-overview/template-overview.component';
import { TemplateViewerComponent } from './templates/template-viewer/template-viewer.component';
import { TournamentLobbyComponent } from './voting/tournament-lobby/tournament-lobby.component';

@NgModule({
  declarations: [
    AppComponent,
    TemplateEditorComponent,
    TemplateListerComponent,
    OptionEditorComponent,
    TemplateOverviewComponent,
    TemplateViewerComponent,
    TournamentLobbyComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    FormsModule,
    BrowserAnimationsModule,
    MatCardModule,
    MatFormFieldModule,
    FormsModule,
    MatListModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
