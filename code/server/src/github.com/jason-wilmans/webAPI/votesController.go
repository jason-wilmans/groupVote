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
	"github.com/satori/go.uuid"
	"fmt"
	domainObjects2 "github.com/jason-wilmans/infrastructure/domainObjects"
)

type VotesController struct {
	voteComponent                *votesComponent.VoteComponent
}

func NewVoteController(voteComponent *votesComponent.VoteComponent) *VotesController {
	precond.NotNil(voteComponent, "")

	controller := &VotesController{voteComponent}
	configureRoutes(controller)
	return controller
}

func (this *VotesController) GetAllTemplates(writer http.ResponseWriter, request *http.Request) {
	templates := this.voteComponent.GetAllTemplates()
	log.Println("Currently has ", len(templates), " templates")
	json.NewEncoder(writer).Encode(templates)
}

func (this *VotesController) CreateTemplate(writer http.ResponseWriter, request *http.Request) {
	var vote domainObjects.Template
	err := json.NewDecoder(request.Body).Decode(&vote)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	this.voteComponent.Create(&vote)
	writer.WriteHeader(http.StatusOK)
}

func (this *VotesController) GetTemplate(writer http.ResponseWriter, request *http.Request) {
	idString := mux.Vars(request)["id"]

	id, err := uuid.FromString(idString)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, err)
		return
	}

	vote := this.voteComponent.GetTemplate(id)

	err = json.NewEncoder(writer).Encode(vote)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (this *VotesController) AddOption(writer http.ResponseWriter, request *http.Request) {
	var option domainObjects.Option
	err := json.NewDecoder(request.Body).Decode(&option)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, err)
		return
	}

	this.voteComponent.AddOption(&option)
	writer.WriteHeader(http.StatusOK)
}

func (this *VotesController) GetOptionsForTemplate(writer http.ResponseWriter, request *http.Request) {
	idString := mux.Vars(request)["id"]

	id, err := uuid.FromString(idString)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, err)
		return
	}

	options := this.voteComponent.GetAllOptionsFor(id)
	json.NewEncoder(writer).Encode(options)
	writer.WriteHeader(http.StatusOK)
}

func (this *VotesController) CreateTournamentFromTemplate(writer http.ResponseWriter, request *http.Request) {

	id, err := routeRegistry.ReadIdParam("id", writer, request)
	if err != nil {
		return
	}

	if !this.voteComponent.TemplateExists(id) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	tournament := this.voteComponent.CreateTournamentFromTemplate(id)
	json.NewEncoder(writer).Encode(tournament)
	writer.WriteHeader(http.StatusOK)
}

func (this *VotesController) GetTournament(writer http.ResponseWriter, request *http.Request) {
	tournamentId, err := routeRegistry.ReadIdParam("tournamentId", writer, request)
	if err != nil {
		return
	}

	tournament := this.voteComponent.GetTournament(tournamentId)
	json.NewEncoder(writer).Encode(tournament)
	writer.WriteHeader(http.StatusOK)
}

func configureRoutes(controller *VotesController) {
	routeRegistry.Register("/templates", routeRegistry.GET, controller.GetAllTemplates)
	routeRegistry.Register("/templates", routeRegistry.POST, controller.CreateTemplate)
	routeRegistry.Register("/templates/{id:" + domainObjects2.EntityIdRegex+ "}", routeRegistry.GET, controller.GetTemplate)
	routeRegistry.Register("/templates/{id:" + domainObjects2.EntityIdRegex+ "}/options", routeRegistry.PUT, controller.AddOption)
	routeRegistry.Register("/templates/{id:" + domainObjects2.EntityIdRegex+ "}/options", routeRegistry.GET, controller.GetOptionsForTemplate)
	routeRegistry.Register("/templates/{id:" + domainObjects2.EntityIdRegex+ "}/tournaments", routeRegistry.POST, controller.CreateTournamentFromTemplate)
	routeRegistry.Register("/tournaments/{tournamentId:" + domainObjects2.EntityIdRegex+ "}", routeRegistry.GET, controller.GetTournament)
}