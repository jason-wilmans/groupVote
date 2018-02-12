package entities

import (
	"github.com/satori/go.uuid"
	"github.com/muroc/prego"
)

type Vote struct {
	Id   uuid.UUID
	Name string
}

func New(name string) Vote {
	id, _ := uuid.NewV4()
	vote := Vote{id, name}

	precond.True(vote.IsValid(), "")
	return vote
}

func (this *Vote) GetId() uuid.UUID {
	return this.Id
}

func (this *Vote) SetName(name string) {
	precond.NotNil(name, "")
	precond.True(len(name) > 0, "")

	this.Name = name
}

//func (this *Vote) Name() string {
//	return this.Name
//}

func (this *Vote) IsValid() bool {
	return len(this.Name) > 3
}