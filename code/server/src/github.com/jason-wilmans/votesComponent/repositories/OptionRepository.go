package repositories

import (
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/satori/go.uuid"
)

type OptionRepository struct {
	options []*domainObjects.Option
}

func NewOptionRepository() *OptionRepository {
	var options []*domainObjects.Option
	return &OptionRepository{options}
}

func (this *OptionRepository) Save(option *domainObjects.Option) {
	precond.NotNil(option, "")

	this.options = append(this.options, option)
}

func (this *OptionRepository) GetAllByVoteId(voteId uuid.UUID) []*domainObjects.Option {
	precond.NotNil(voteId, "")

	var filteredOptions []*domainObjects.Option
	for _, option := range this.options {
		if uuid.Equal(option.GetVoteId(), voteId) {
			filteredOptions = append(filteredOptions, option)
		}
	}

	return filteredOptions
}

func (this *OptionRepository) Exists(id uuid.UUID) bool {
	for _, option := range this.options {
		if option.GetId() == id {
			return true
		}
	}

	return false
}