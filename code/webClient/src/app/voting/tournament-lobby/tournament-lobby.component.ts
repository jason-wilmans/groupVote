import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {VoteService} from "../../model/vote-service";
import {ClientDataService} from "../../infrastructure/client-data.service";

@Component({
  selector: 'app-tournament-lobby',
  templateUrl: './tournament-lobby.component.html',
  styleUrls: ['./tournament-lobby.component.css']
})
export class TournamentLobbyComponent implements OnInit {
  private TournamentId: string;
  private TemplateId: string;
  private tournament;
  private AlreadyHasId: boolean;

  constructor(
    private voteService: VoteService,
    private clientDataService: ClientDataService,
    private route: ActivatedRoute
  ) { }

  async ngOnInit() {
    this.TournamentId = this.route.snapshot.paramMap.get('id');

    if (this.TournamentId) {
      this.tournament = await this.voteService.GetTournament(this.TournamentId);
      const template = await this.voteService.Get(this.tournament.TemplateId);
      this.Name = template.Name;
    }
  }

  EnterTournament() {
    this.voteService.EnterTournament(this.TournamentId)
  }

}
