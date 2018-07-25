import { Component, OnInit } from '@angular/core';
import {VoteService} from "../../model/vote-service";
import {Template} from "../../model/template";

@Component({
  selector: 'app-vote-lister',
  templateUrl: './vote-lister.component.html',
  styleUrls: ['./vote-lister.component.css']
})
export class VoteListerComponent implements OnInit {
  public Templates: Template[];

  constructor(
    private voteService: VoteService
  ) {
    voteService.VotesChanged.subscribe(test => {
      this.ngOnInit();
    })
  }

  async ngOnInit() {
    this.Templates = await this.voteService.GetAllTemplates();
  }

}
