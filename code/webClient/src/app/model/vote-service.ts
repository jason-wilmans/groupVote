import {EventEmitter, Injectable} from "@angular/core";
import {Template} from "./template";
import {HttpClient} from "@angular/common/http";
import {Option} from "./option";

@Injectable({
  providedIn: 'root'
})
export class VoteService {
  public VotesChanged : EventEmitter<any> = new EventEmitter();

  constructor(
    private http: HttpClient
  ) {}

  public async Create(template: Template) {
    await this.http
      .post<Template>("http://localhost:8080/templates", template)
      .toPromise();

    this.VotesChanged.emit();
  }

  public async GetAllTemplates() : Promise<Template[]> {
    return this.http
      .get<Template[]>("http://localhost:8080/templates")
      .toPromise();
  }

  public async Get(id: string) : Promise<Template> {
    return this.http
      .get<Template>("http://localhost:8080/templates/" + id)
      .toPromise();
  }

  async Save(vote: Template) : Promise<void> {
    return this.http
      .post<void>("http://localhost:8080/templates/" + vote.GetId(), vote)
      .toPromise();
  }

  public AddOption(option: Option) : Promise<void> {
    return this.http
      .put<void>("http://localhost:8080/templates/" + option.GetTemplateId() + "/options", option)
      .toPromise();
  }

  async GetAllOptions(templateId: string) : Promise<Option[]> {
    return this.http
      .get<Option[]>("http://localhost:8080/templates/" + templateId + "/options")
      .toPromise();
  }
}
