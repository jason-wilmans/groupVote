package domainObjects

import (
	"github.com/satori/go.uuid"
	"github.com/muroc/prego"
)

type Template struct {
	Id   uuid.UUID
	Name string
}

func NewTemplate(name string) *Template {
	precond.True(len(name) > 0, "")

	id, _ := uuid.NewV4()
	return &Template{id, name}
}

func (this *Template) GetId() uuid.UUID {
	return this.Id
}

func (this *Template) SetName(name string) {
	precond.NotNil(name, "")
	precond.True(len(name) > 0, "")

	this.Name = name
}

func (this *Template) GetName() string {
	return this.Name
}