import { Component, OnInit } from '@angular/core';
import {VoteService} from "../vote.service";

@Component({
  selector: 'app-votes',
  templateUrl: './votes.component.html',
  styleUrls: ['./votes.component.css']
})

export class VotesComponent implements OnInit {

  constructor(private votesService : VoteService) { }

  ngOnInit() {
  }

}
