package votesComponent

import (
	"github.com/jason-wilmans/votesComponent/useCases"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/satori/go.uuid"
)

type VoteComponent struct {
	voteUC *useCases.VoteUC
	optionUC *useCases.OptionUC
}

func New() *VoteComponent {
	votesRepository := repositories.NewVoteRepository()
	optionRepository := repositories.NewOptionRepository()

	voteUC := useCases.NewVoteUC(votesRepository)
	optionUC := useCases.NewOptionUC(optionRepository, voteUC)

	return &VoteComponent{voteUC, optionUC}
}

func (this *VoteComponent) Create(vote *domainObjects.Vote) {
	this.voteUC.Create(vote)
}

func (this *VoteComponent) GetAll() []*domainObjects.Vote {
	return this.voteUC.GetAll()
}

func (this *VoteComponent) AddOption(option *domainObjects.Option) {
	this.optionUC.AddOption(option)
}
func (this *VoteComponent) GetVote(id uuid.UUID) *domainObjects.Vote {
	return this.voteUC.Get(id)
}