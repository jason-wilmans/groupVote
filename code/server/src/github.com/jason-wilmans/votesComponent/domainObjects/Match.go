package domainObjects

import (
	"github.com/satori/go.uuid"
	"github.com/muroc/prego"
)

type Match struct {
	Id				string
	CompetitorAId	string
	CompetitorBId	string
	Started			bool
	Finished		bool
}

func NewMatch() *Match {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return &Match{id.String(), "", "", false, false}
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