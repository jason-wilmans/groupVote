import { UUID } from "angular2-uuid";

export class Vote {
  private Id : string;
  private Title : string;

  constructor(title : string) {
    this.Id = UUID.UUID();
    this.Title = title;
  }

  getId(): string {
    return this.Id;
  }

  get getTitle(): string {
    return this.Title;
  }

  setTitle(value: string) {
    this.Title = value;
  }
}
