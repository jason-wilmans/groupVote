import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { VotesComponent } from './votes/votes.component';
import {VoteService} from "./vote.service";
import {HttpClient, HttpHandler, HttpClientModule} from "@angular/common/http";

@NgModule({
  declarations: [
    AppComponent,
    VotesComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [VoteService, HttpClient, HttpHandler],
  bootstrap: [AppComponent]
})

export class AppModule {

}
