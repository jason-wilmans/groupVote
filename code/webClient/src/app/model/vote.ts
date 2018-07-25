import {UUID} from "uuid";
import {Option} from "./option";

export class Vote {
  private readonly Id: string;
  public Name: string;
  private Options: Option[];

  constructor(name: string, id?: string) {

    if (id) {
      this.Id = id;
    } else {

    }

    this.Name = name;
    this.Options = [];
  }

  public GetId(): string {
    return this.Id;
  }

  public SetName(name: string) {
    this.Name = name;
  }

  public GetName(): string {
    return this.Name;
  }

  public Add(option: Option) {
    this.Options.push(option)
  }

  GetOptions(): Option[] {
    const result: Option[] = [];
    this.Options.forEach((option) => {
      result.push(option)
    });
    return result;
  }

}
