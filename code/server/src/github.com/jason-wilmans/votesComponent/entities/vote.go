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
	precond.True(len(name) > 0, "")

	id, _ := uuid.NewV4()
	return Vote{id, name}
}

func (this *Vote) GetId() uuid.UUID {
	return this.Id
}

func (this *Vote) SetName(name string) {
	precond.NotNil(name, "")
	precond.True(len(name) > 0, "")

	this.Name = name
}

func (this *Vote) GetName() string {
	return this.Name
}