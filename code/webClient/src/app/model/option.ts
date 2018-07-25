export class Option {
  private readonly Id: string;
  private readonly TemplateId: string;
  private Name: string;
  private Description: string;

  constructor(templateId: string, name: string, description?: string, id?: string) {
    this.TemplateId = templateId;
    this.Name = name;
    this.Description = description;
    if(id) {
      this.Id = id;
    }
  }

  public GetId() : string {
    return this.Id;
  }

  public GetTemplateId(): string {
    return this.TemplateId;
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
