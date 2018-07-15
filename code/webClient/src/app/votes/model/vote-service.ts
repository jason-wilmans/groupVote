import {Injectable} from "@angular/core";
import {Vote} from "./vote";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class VoteService {
  constructor(
    private http: HttpClient
  ) {}

  public async Create(vote: Vote) {
    return this.http
      .post<Vote>("http://localhost:8080/votes", vote)
      .toPromise();
  }
}
