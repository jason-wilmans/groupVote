package domainObjects

import (
	"github.com/satori/go.uuid"
	"github.com/muroc/prego"
)

type Match struct {
	Id				string
	TournamentId	string
	Level			int
	Index			int
	CompetitorAId	string
	CompetitorBId	string
	Started			bool
	Finished		bool
}

func NewMatch(tournamentId uuid.UUID, level int, index int) *Match {
	precond.NotNil(tournamentId, "")
	precond.True(level >= 0, "")
	precond.True(index >= 0, "")

	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return &Match{id.String(), tournamentId.String(), level, index,"", "", false, false}
}

func (this *Match) GetId() uuid.UUID {
	id, err := uuid.FromString(this.Id)
	if err != nil {
		panic(err)
	}
	return id
}

func (this *Match) HasRoom() bool {
	return this.CompetitorAId == "" || this.CompetitorBId == ""
}

func (this *Match) Join(competitorId uuid.UUID) {
	precond.True(this.HasRoom(), "")
	precond.False(this.Started || this.Finished, "")

	if this.CompetitorAId == "" {
		this.CompetitorAId = competitorId.String()
		return
	}

	this.CompetitorBId = competitorId.String()
}

func (this *Match) GetA() string {
	return this.CompetitorAId
}

func (this *Match) GetB() string {
	return this.CompetitorBId
}