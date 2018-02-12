package votesController

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/jason-wilmans/votesComponent/entities"
	"github.com/jason-wilmans/votesComponent/repositories"
	"github.com/muroc/prego"
)

type VotesController struct {
	repository *repositories.VotesRepository
}

func New(repository *repositories.VotesRepository) *VotesController {
	precond.NotNil(repository, "")

	return &VotesController{repository}
}

func (this *VotesController) GetAllVotes(writer http.ResponseWriter, request *http.Request) {
	log.Println("All votes requested.")

	//votes := this.repository.GetAll()

	//log.Println("Currently has ", len(votes), " votes")

	json.NewEncoder(writer).Encode(entities.New("Test?"))
}

func (this *VotesController) CreateVote(writer http.ResponseWriter, request *http.Request)  {
	log.Println("Create a vote requested.")

	var vote entities.Vote
	_ = json.NewDecoder(request.Body).Decode(&vote)

	if !vote.IsValid() {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Invalid values"))
	}

	if this.repository.Exists(vote.GetId()) {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Already exists"))
	}

	this.repository.AddNew(vote)
	writer.WriteHeader(200)
}