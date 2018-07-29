package repositories

import (
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type VoteRepository struct {
	matches []*domainObjects.Template
}

func NewVoteRepository() *VoteRepository {
	var votes []*domainObjects.Template
	return &VoteRepository{votes}
}

func (this *VoteRepository) Save(vote *domainObjects.Template) {
	precond.NotNil(vote, "")

	this.matches = append(this.matches, vote)
}

func (this *VoteRepository) GetAll() []*domainObjects.Template {
	return this.matches
}

func (this *VoteRepository) Exists(id uuid.UUID) bool {
	for _, vote := range this.matches {
		if vote.GetId() == id {
			return true
		}
	}

	return false
}
func (this *VoteRepository) Get(id uuid.UUID) *domainObjects.Template {
	for _, vote := range this.matches {
		if vote.GetId() == id {
			return vote
		}
	}

	return nil
}