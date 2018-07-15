package votesController

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/jason-wilmans/votesComponent/entities"
	"github.com/muroc/prego"
	"github.com/jason-wilmans/votesComponent/service"
)

type VotesController struct {
	service *service.VoteService
}

func New(service *service.VoteService) *VotesController {
	precond.NotNil(service, "")

	return &VotesController{service}
}

func (this *VotesController) GetAllVotes(writer http.ResponseWriter, request *http.Request) {
	log.Println("All votes requested.")

	votes := this.service.GetAll()
	log.Println("Currently has ", len(votes), " votes")
	json.NewEncoder(writer).Encode(votes)
}

func (this *VotesController) CreateVote(writer http.ResponseWriter, request *http.Request)  {
	log.Println("Save a vote requested.")

	var vote entities.Vote
	_ = json.NewDecoder(request.Body).Decode(&vote)

	this.service.Create(vote)
	writer.WriteHeader(200)
}