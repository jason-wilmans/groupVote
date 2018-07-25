package useCases

import (
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type OptionUC struct {
	repository *repositories.OptionRepository
	voteUC *VoteUC
}

func NewOptionUC(repository *repositories.OptionRepository, voteUC *VoteUC) *OptionUC {
	precond.NotNil(repository, "")
	precond.NotNil(voteUC, "")

	return &OptionUC{repository, voteUC}
}

func (this *OptionUC) AddOption(toAdd *domainObjects.Option) {
	precond.NotNil(toAdd, "")
	precond.False(this.repository.Exists(toAdd.GetId()), "")
	precond.True(this.voteUC.Exists(toAdd.GetVoteId()), "")

	option := domainObjects.NewOption(toAdd.VoteId, toAdd.GetName(), toAdd.GetDescription())
	this.repository.Save(option)
}

func (this *OptionUC) GetAllForVoteId(voteId uuid.UUID) []*domainObjects.Option  {
	precond.True(this.voteUC.Exists(voteId), "")

	return this.repository.GetAllByVoteId(voteId)
}