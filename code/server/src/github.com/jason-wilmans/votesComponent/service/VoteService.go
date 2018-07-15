package service

import (
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/jason-wilmans/votesComponent/entities"
	"github.com/muroc/prego"
)

type VoteService struct {
	repository *repositories.VotesRepository
}

func New(repository *repositories.VotesRepository) *VoteService {
	return &VoteService{repository}
}

func (this VoteService) GetAll() []entities.Vote {
	return this.repository.GetAll()
}
func (this VoteService) Create(toCreate entities.Vote) {
	precond.False(this.repository.Exists(toCreate.Id), "Vote with this id already exists.")

	vote := entities.New(toCreate.GetName())
	this.repository.Save(vote)
}

