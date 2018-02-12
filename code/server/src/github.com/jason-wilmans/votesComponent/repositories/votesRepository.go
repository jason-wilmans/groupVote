package repositories

import (
	"github.com/jason-wilmans/votesComponent/entities"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type VotesRepository struct {
	votes []entities.Vote
}

func New() *VotesRepository {
	var votes []entities.Vote
	vote := entities.New("Test Vote")
	votes = append(votes, vote)
	return &VotesRepository{votes}
}

func (this *VotesRepository) AddNew(vote entities.Vote) {
	precond.NotNil(vote, "")
	precond.True(vote.IsValid(), "")

	this.votes = append(this.votes, vote)
}

func (this *VotesRepository) GetAll() []entities.Vote {
	return this.votes
}

func (this *VotesRepository) Exists(id uuid.UUID) bool {
	for _, vote := range this.votes {
		if vote.GetId() == id {
			return true
		}
	}

	return false
}