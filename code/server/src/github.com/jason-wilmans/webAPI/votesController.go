package webAPI

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/jason-wilmans/infrastructure/routeRegistry"
	"github.com/jason-wilmans/votesComponent"
)

type VotesController struct {
	voteComponent *votesComponent.VotesComponent
}

func NewVoteController(voteComponent *votesComponent.VotesComponent) *VotesController {
	precond.NotNil(voteComponent, "")

	controller := &VotesController{voteComponent}
	configureRoutes(controller)
	return controller
}

func (this *VotesController) GetAllVotes(writer http.ResponseWriter, request *http.Request) {
	log.Println("All votes requested.")

	votes := this.voteComponent.GetAll()
	log.Println("Currently has ", len(votes), " votes")
	json.NewEncoder(writer).Encode(votes)
}

func (this *VotesController) CreateVote(writer http.ResponseWriter, request *http.Request)  {
	log.Println("Save a vote requested.")

	var vote domainObjects.Vote
	_ = json.NewDecoder(request.Body).Decode(&vote)

	this.voteComponent.Create(&vote)
	writer.WriteHeader(200)
}

func configureRoutes(controller *VotesController) {
	routeRegistry.Register("/votes", routeRegistry.GET, controller.GetAllVotes)
	routeRegistry.Register("/votes", routeRegistry.POST, controller.CreateVote)
}