import { Component, OnInit } from '@angular/core';
import {VoteService} from "../../model/vote-service";
import {Template} from "../../model/template";

@Component({
  selector: 'app-template-lister',
  templateUrl: './template-lister.component.html',
  styleUrls: ['./template-lister.component.css']
})
export class TemplateListerComponent implements OnInit {
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
