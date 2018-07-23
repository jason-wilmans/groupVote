import { Component, OnInit } from '@angular/core';
import {VoteService} from "../../model/vote-service";
import {Vote} from "../../model/vote";

@Component({
  selector: 'app-vote-lister',
  templateUrl: './vote-lister.component.html',
  styleUrls: ['./vote-lister.component.css']
})
export class VoteListerComponent implements OnInit {
  public Votes : Vote[];

  constructor(
    private voteService: VoteService
  ) {
    voteService.VotesChanged.subscribe(test => {
      this.ngOnInit();
    })
  }

  async ngOnInit() {
    this.Votes = await this.voteService.GetAllVotes();
  }

}
