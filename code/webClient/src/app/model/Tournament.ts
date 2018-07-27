import {Match} from "./Match";

export class Tournament {
  private Id: string ;
  private Matches: Match[];

  constructor(
  ) {
  }

  public GetId() : string {
    return this.Id;
  }
}
