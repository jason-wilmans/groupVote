import { Component, OnInit } from '@angular/core';
import {Option} from "../../model/option";
import {VoteService} from "../../model/vote-service";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-template-viewer',
  templateUrl: './template-viewer.component.html',
  styleUrls: ['./template-viewer.component.css']
})
export class TemplateViewerComponent implements OnInit {

  private TemplateId: string | null;
  private Name: string;
  private Options: Option[];

  constructor(
    private voteService: VoteService,
    private route: ActivatedRoute
  ) { }

  async ngOnInit() {
    this.TemplateId = this.route.snapshot.paramMap.get('id');

    if (this.TemplateId) {
      const vote = await this.voteService.Get(this.TemplateId);
      this.Name = vote.Name;
      this.Options = await this.voteService.GetAllOptions(this.TemplateId)
    }
  }
}
