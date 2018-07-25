import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {VoteService} from "../../model/vote-service";
import {Option} from "../../model/option";

@Component({
  selector: 'app-option-editor',
  templateUrl: './option-editor.component.html',
  styleUrls: ['./option-editor.component.css']
})
export class OptionEditorComponent implements OnInit {

  private Name: string;
  private Description: string;

  constructor(
    private voteService: VoteService,
    private route: ActivatedRoute
  ) { }

  async OnAddOptionClicked() {
    let voteId = this.route.snapshot.paramMap.get('id');
    const option: Option = new Option(voteId, this.Name, this.Description);
    console.log("sending: ", option);
    await this.voteService.AddOption(option);
  }

  ngOnInit() {
  }

}
