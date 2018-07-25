import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
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
  private TemplateId: string | null;

  constructor(
    private router: Router,
    private voteService: VoteService,
    private route: ActivatedRoute
  ) { }

  async OnAddOptionClicked() {
    const option: Option = new Option(this.TemplateId, this.Name, this.Description);
    await this.voteService.AddOption(option);
    await this.router.navigateByUrl("/templates/" + this.TemplateId);
  }

  async ngOnInit() {
    this.TemplateId = this.route.snapshot.paramMap.get('id');
  }

}
