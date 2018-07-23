import {Component} from '@angular/core';
import {VoteService} from "../../model/vote-service";
import {Vote} from "../../model/vote";
import {contract} from "typedcontract";

@Component({
  selector: 'app-vote-editor',
  templateUrl: './vote-editor.component.html',
  styleUrls: ['./vote-editor.component.css']
})
export class VoteEditorComponent {

  public Name: string;
  private CanCreate: boolean;

  constructor(
    private voteService: VoteService
  ) {
    this.CanCreate = true;
    this.Name = "";
  }

  public async OnCreateVoteClicked() {
    contract.In(this.CanCreate).isTrue();

    const vote = new Vote(this.Name);
    await this.voteService.Create(vote);
    this.Name = '';
  }

}
