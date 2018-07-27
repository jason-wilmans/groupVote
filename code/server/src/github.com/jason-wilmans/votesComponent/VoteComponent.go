package votesComponent

import (
	"github.com/jason-wilmans/votesComponent/useCases"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/satori/go.uuid"
)

type VoteComponent struct {
	templateUC   *useCases.TemplateUC
	optionUC     *useCases.OptionUC
	tournamentUC *useCases.TournamentUC
}

func New() *VoteComponent {
	votesRepository := repositories.NewVoteRepository()
	optionRepository := repositories.NewOptionRepository()

	voteUC := useCases.NewVoteUC(votesRepository)
	optionUC := useCases.NewOptionUC(optionRepository, voteUC)
	tournamentUC := useCases.NewTournamentUC(optionUC)

	return &VoteComponent{voteUC, optionUC, tournamentUC}
}

func (this *VoteComponent) Create(vote *domainObjects.Template) {
	this.templateUC.Create(vote)
}

func (this *VoteComponent) GetTemplate(id uuid.UUID) *domainObjects.Template {
	return this.templateUC.Get(id)
}

func (this *VoteComponent) GetAllTemplates() []*domainObjects.Template {
	return this.templateUC.GetAll()
}

func (this *VoteComponent) AddOption(option *domainObjects.Option) {
	this.optionUC.AddOption(option)
}

func (this *VoteComponent) GetAllOptionsFor(voteId uuid.UUID) []*domainObjects.Option {
	return this.optionUC.GetAllForTemplateId(voteId)
}
func (this *VoteComponent) CreateTournamentFromTemplate(templateId uuid.UUID) *domainObjects.Tournament {
	return this.tournamentUC.CreateTournament(templateId)
}
func (this *VoteComponent) GetTournament(tournamentId uuid.UUID) *domainObjects.Tournament {
	return this.tournamentUC.GetTournament(tournamentId)
}
func (this *VoteComponent) TemplateExists(templateId uuid.UUID) bool {
	return this.templateUC.Exists(templateId)
}