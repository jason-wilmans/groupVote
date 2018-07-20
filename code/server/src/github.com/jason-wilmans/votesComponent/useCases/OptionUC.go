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

func (this *OptionUC) AddOption(option *domainObjects.Option) {
	precond.NotNil(option, "")
	precond.False(this.repository.Exists(option.Id), "")
	precond.True(this.voteUC.Exists(option.VoteId), "")

	this.repository.Save(option)
}

func (this *OptionUC) GetAllByVoteId(voteId uuid.UUID) []*domainObjects.Option  {
	precond.True(this.voteUC.Exists(voteId), "")

	return this.repository.GetAllByVoteId(voteId)
}