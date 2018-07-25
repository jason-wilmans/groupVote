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
  private VoteId: string | null;


  constructor(
    private voteService: VoteService,
    private route: ActivatedRoute
  ) { }

  async OnAddOptionClicked() {
    const option: Option = new Option(this.VoteId, this.Name, this.Description);
    await this.voteService.AddOption(option);
  }

  async ngOnInit() {
    this.VoteId = this.route.snapshot.paramMap.get('id');
  }

}
