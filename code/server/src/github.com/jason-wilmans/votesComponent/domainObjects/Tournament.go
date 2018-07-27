package domainObjects

import (
	"github.com/satori/go.uuid"
	"github.com/muroc/prego"
	)

type Tournament struct {
	Id		string
	Matches []*Match
}

func NewTournament(matches []*Match) *Tournament {
	precond.NotNil(matches, "")
	precond.True(len(matches) > 0, "")

	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	idString := id.String()
	return &Tournament{idString, matches}
}

func (this *Tournament) GetId() uuid.UUID {
	id, err := uuid.FromString(this.Id)
	if err != nil {
		panic(err)
	}
	return id
}