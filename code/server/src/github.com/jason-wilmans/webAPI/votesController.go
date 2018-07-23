package webAPI

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"github.com/muroc/prego"
	"github.com/jason-wilmans/infrastructure/routeRegistry"
	"github.com/jason-wilmans/votesComponent"
	"github.com/gorilla/mux"
	"github.com/jason-wilmans/infrastructure/goExtensions"
)

type VotesController struct {
	voteComponent *votesComponent.VoteComponent
}

func NewVoteController(voteComponent *votesComponent.VoteComponent) *VotesController {
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

func (this *VotesController) CreateVote(writer http.ResponseWriter, request *http.Request) {
	log.Println("Save a vote requested.")

	var vote domainObjects.Vote
	err := json.NewDecoder(request.Body).Decode(&vote)

	if err != nil {
		writer.WriteHeader(400)
		return
	}

	this.voteComponent.Create(&vote)
	writer.WriteHeader(200)
}

func (this *VotesController) AddOption(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]

	log.Println("AddOption called with " + id)

	var option domainObjects.Option
	err := json.NewDecoder(request.Body).Decode(&option)

	if err != nil {
		writer.WriteHeader(400)
		return
	}

	this.voteComponent.AddOption(&option)
	writer.WriteHeader(200)
}

func configureRoutes(controller *VotesController) {
	routeRegistry.Register("/votes", routeRegistry.GET, controller.GetAllVotes)
	routeRegistry.Register("/votes", routeRegistry.POST, controller.CreateVote)
	routeRegistry.Register("/votes/{id:" + goExtensions.UUIDRegEx + "}", routeRegistry.POST, controller.AddOption)
}