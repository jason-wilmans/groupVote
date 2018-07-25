package repositories

import (
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type VoteRepository struct {
	votes []*domainObjects.Template
}

func NewVoteRepository() *VoteRepository {
	var votes []*domainObjects.Template
	return &VoteRepository{votes}
}

func (this *VoteRepository) Save(vote *domainObjects.Template) {
	precond.NotNil(vote, "")

	this.votes = append(this.votes, vote)
}

func (this *VoteRepository) GetAll() []*domainObjects.Template {
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
func (this *VoteRepository) Get(id uuid.UUID) *domainObjects.Template {
	for _, vote := range this.votes {
		if vote.GetId() == id {
			return vote
		}
	}

	return nil
}