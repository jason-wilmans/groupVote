package votesComponent

import (
	"github.com/jason-wilmans/votesComponent/controller"
	"github.com/jason-wilmans/infrastructure/routeRegistry"
	"github.com/jason-wilmans/votesComponent/repositories"
)

type VotesComponent struct {
	controller *votesController.VotesController
	service		*repositories.VotesRepository
}

func New() *VotesComponent {
	service := repositories.New()
	controller := votesController.New(service)

	routeRegistry.Register("/votes", routeRegistry.GET, controller.GetAllVotes)
	routeRegistry.Register("/votes", routeRegistry.POST, controller.CreateVote)

	return &VotesComponent{controller, service}
}