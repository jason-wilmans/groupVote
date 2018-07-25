export class Option {
  private readonly Id: string;
  private readonly VoteId: string;
  private Name: string;
  private Description: string;

  constructor(voteId: string, name: string, description?: string, id?: string) {
    this.VoteId = voteId;
    this.Name = name;
    this.Description = description;
    if(id) {
      this.Id = id;
    }
  }

  public GetId() : string {
    return this.Id;
  }

  public GetVoteId(): string {
    return this.VoteId;
  }

  public SetName(name: string) {
    this.Name = name;
  }

  public GetName() : string {
    return this.Name;
  }

  public SetDescription(description: string) {
    this.Description = description;
  }

  public GetDescription() : string {
    return this.Description;
  }
}
