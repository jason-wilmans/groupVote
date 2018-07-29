package domainObjects

import (
	"github.com/satori/go.uuid"
	"github.com/muroc/prego"
)

type Tournament struct {
	Id			string
	TemplateId	string
	Voters		[]uuid.UUID
}

func NewTournament(templateId uuid.UUID) *Tournament {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	voters := make([]uuid.UUID, 0)

	return &Tournament{id.String(), templateId.String(), voters}
}

func (this *Tournament) GetId() uuid.UUID {
	id, err := uuid.FromString(this.Id)
	if err != nil {
		panic(err)
	}
	return id
}

func (this *Tournament) AddVoter(voterId uuid.UUID) {
	precond.False(this.HasVoter(voterId), "")

	this.Voters = append(this.Voters, voterId)
}

func (this *Tournament) HasVoter(voterId uuid.UUID) bool {
	precond.NotNil(voterId, "")

	for _, existing := range this.Voters {
		if uuid.Equal(voterId, existing) {
			return true
		}
	}
	return false
}