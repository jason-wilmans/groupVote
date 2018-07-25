import { Component, OnInit } from '@angular/core';
import {Option} from "../../model/option";
import {VoteService} from "../../model/vote-service";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-vote-viewer',
  templateUrl: './vote-viewer.component.html',
  styleUrls: ['./vote-viewer.component.css']
})
export class VoteViewerComponent implements OnInit {

  private VoteId: string | null;
  private Name: string;
  private Options: Option[];

  constructor(
    private voteService: VoteService,
    private route: ActivatedRoute
  ) { }

  async ngOnInit() {
    this.VoteId = this.route.snapshot.paramMap.get('id');

    if (this.VoteId) {
      const vote = await this.voteService.Get(this.VoteId);
      this.Name = vote.Name;
      this.Options = await this.voteService.GetAllOptions(this.VoteId)
    }
  }

}
