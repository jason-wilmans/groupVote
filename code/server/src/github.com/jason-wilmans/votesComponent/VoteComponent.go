package votesComponent

import (
	"github.com/jason-wilmans/votesComponent/useCases"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/satori/go.uuid"
)

type VoteComponent struct {
	voteUC *useCases.TemplateUC
	optionUC *useCases.OptionUC
}

func New() *VoteComponent {
	votesRepository := repositories.NewVoteRepository()
	optionRepository := repositories.NewOptionRepository()

	voteUC := useCases.NewVoteUC(votesRepository)
	optionUC := useCases.NewOptionUC(optionRepository, voteUC)

	return &VoteComponent{voteUC, optionUC}
}

func (this *VoteComponent) Create(vote *domainObjects.Template) {
	this.voteUC.Create(vote)
}

func (this *VoteComponent) GetTemplate(id uuid.UUID) *domainObjects.Template {
	return this.voteUC.Get(id)
}

func (this *VoteComponent) GetAllTemplates() []*domainObjects.Template {
	return this.voteUC.GetAll()
}

func (this *VoteComponent) AddOption(option *domainObjects.Option) {
	this.optionUC.AddOption(option)
}

func (this *VoteComponent) GetAllOptionsFor(voteId uuid.UUID) []*domainObjects.Option {
	return this.optionUC.GetAllForVoteId(voteId)
}