package useCases

import (
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type TemplateUC struct {
	repository *repositories.VoteRepository
}

func NewVoteUC(repository *repositories.VoteRepository) *TemplateUC {
	return &TemplateUC{repository}
}

func (this TemplateUC) GetAll() []*domainObjects.Template {
	return this.repository.GetAll()
}

func (this TemplateUC) Create(toCreate *domainObjects.Template) {
	precond.False(this.repository.Exists(toCreate.Id), "Template with this id already exists.")

	vote := domainObjects.NewTemplate(toCreate.GetName())
	this.repository.Save(vote)
}

func (this *TemplateUC) Exists(templateId uuid.UUID) bool {
	return this.repository.Exists(templateId)
}
func (this *TemplateUC) Get(id uuid.UUID) *domainObjects.Template {
	return this.repository.Get(id)
}

