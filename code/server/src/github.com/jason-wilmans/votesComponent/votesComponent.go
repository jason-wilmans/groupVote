package votesComponent

import (
	"github.com/jason-wilmans/votesComponent/controller"
	"github.com/jason-wilmans/infrastructure/routeRegistry"
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/jason-wilmans/votesComponent/service"
)

type VotesComponent struct {
	controller *votesController.VotesController
	service *service.VoteService
}

func New() *VotesComponent {
	repository := repositories.New()
	voteService := service.New(repository)
	controller := votesController.New(voteService)

	routeRegistry.Register("/votes", routeRegistry.GET, controller.GetAllVotes)
	routeRegistry.Register("/votes", routeRegistry.POST, controller.CreateVote)

	return &VotesComponent{controller, voteService}
}