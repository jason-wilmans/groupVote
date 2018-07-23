import {EventEmitter, Injectable} from "@angular/core";
import {Vote} from "./vote";
import {HttpClient} from "@angular/common/http";

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
}
