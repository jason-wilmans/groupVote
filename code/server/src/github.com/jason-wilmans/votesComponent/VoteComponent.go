package votesComponent

import (
	"github.com/jason-wilmans/votesComponent/useCases"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/jason-wilmans/votesComponent/repositories"
)

type VotesComponent struct {
	voteUC *useCases.VoteUC
	optionUC *useCases.OptionUC
}

func New() *VotesComponent {
	votesRepository := repositories.NewVoteRepository()
	optionRepository := repositories.NewOptionRepository()

	voteUC := useCases.NewVoteUC(votesRepository)
	optionUC := useCases.NewOptionUC(optionRepository, voteUC)

	return &VotesComponent{voteUC, optionUC}
}

func (this *VotesComponent) Create(vote *domainObjects.Vote) {
	this.voteUC.Create(vote)
}

func (this *VotesComponent) GetAll() []*domainObjects.Vote {
	return this.voteUC.GetAll()
}

