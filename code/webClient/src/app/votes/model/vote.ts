import {UUID} from "uuid";

export class Vote {
  private Id: string;
  private Name: string;

  constructor(name: string, id?: string) {

    if(id) {
      this.Id = id;
    } else {

    }

    this.Name = name;
  }
}
