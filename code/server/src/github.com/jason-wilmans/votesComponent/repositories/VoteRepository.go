package repositories

import (
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type VoteRepository struct {
	votes []*domainObjects.Vote
}

func NewVoteRepository() *VoteRepository {
	var votes []*domainObjects.Vote
	return &VoteRepository{votes}
}

func (this *VoteRepository) Save(vote *domainObjects.Vote) {
	precond.NotNil(vote, "")

	this.votes = append(this.votes, vote)
}

func (this *VoteRepository) GetAll() []*domainObjects.Vote {
	return this.votes
}

func (this *VoteRepository) Exists(id uuid.UUID) bool {
	for _, vote := range this.votes {
		if vote.GetId() == id {
			return true
		}
	}

	return false
}
func (this *VoteRepository) Get(id uuid.UUID) *domainObjects.Vote {
	for _, vote := range this.votes {
		if vote.GetId() == id {
			return vote
		}
	}

	return nil
}