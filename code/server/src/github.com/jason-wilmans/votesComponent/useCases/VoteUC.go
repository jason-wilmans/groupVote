package useCases

import (
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type VoteUC struct {
	repository *repositories.VoteRepository
}

func NewVoteUC(repository *repositories.VoteRepository) *VoteUC {
	return &VoteUC{repository}
}

func (this VoteUC) GetAll() []*domainObjects.Template {
	return this.repository.GetAll()
}

func (this VoteUC) Create(toCreate *domainObjects.Template) {
	precond.False(this.repository.Exists(toCreate.Id), "Template with this id already exists.")

	vote := domainObjects.NewTemplate(toCreate.GetName())
	this.repository.Save(vote)
}

func (this *VoteUC) Exists(voteId uuid.UUID) bool {
	return this.repository.Exists(voteId)
}
func (this *VoteUC) Get(id uuid.UUID) *domainObjects.Template {
	return this.repository.Get(id)
}

