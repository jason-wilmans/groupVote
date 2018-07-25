import {EventEmitter, Injectable} from "@angular/core";
import {Vote} from "./vote";
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

  public async Create(vote: Vote) {
    await this.http
      .post<Vote>("http://localhost:8080/votes", vote)
      .toPromise();

    this.VotesChanged.emit();
  }

  public async GetAllVotes() : Promise<Vote[]> {
    return this.http
      .get<Vote[]>("http://localhost:8080/votes")
      .toPromise();
  }

  public async Get(id: string) : Promise<Vote> {
    return this.http
      .get<Vote>("http://localhost:8080/votes/" + id)
      .toPromise();
  }

  async Save(vote: Vote) : Promise<void> {
    return this.http
      .post<void>("http://localhost:8080/votes/" + vote.GetId(), vote)
      .toPromise();
  }

  public AddOption(option: Option) : Promise<void> {
    return this.http
      .put<void>("http://localhost:8080/votes/" + option.GetVoteId() + "/options", option)
      .toPromise();
  }
}
