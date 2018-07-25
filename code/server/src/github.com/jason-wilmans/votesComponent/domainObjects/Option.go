package domainObjects

import (
	"github.com/satori/go.uuid"
	"github.com/muroc/prego"
	"github.com/jason-wilmans/infrastructure/goExtensions"
)

type Option struct {
	Id			string
	VoteId		string
	Name		string
	Description string
}

func NewOption(voteId string, name string, description string) *Option {
	precond.NotNil(voteId, "")
	precond.NotNil(name, "")
	precond.True(goExtensions.IsNonWhitespaceWithMinimumLength(name, 1), "")

	id, _ := uuid.NewV4()
	return &Option{id.String(), voteId,name,	description}
}

func (this *Option) GetId() uuid.UUID {
	id, _ := uuid.FromString(this.Id)
	return id
}

func (this *Option) SetName(name string) {
	this.Name = name
}

func (this *Option) GetName() string {
	return this.Name
}

func (this *Option) SetDescription(description string) {
	this.Description = description
}

func (this *Option) GetDescription() string {
	return this.Description
}

func (this *Option) GetVoteId() uuid.UUID {
	id, _ := uuid.FromString(this.VoteId)
	return id
}