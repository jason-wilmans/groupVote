package useCases

import (
	"github.com/satori/go.uuid"
	"github.com/jason-wilmans/votesComponent/domainObjects"
)

type TournamentUC struct {
	tournaments		map[uuid.UUID]*domainObjects.Tournament
	optionUC		*OptionUC
}

func NewTournamentUC(optionUC *OptionUC) *TournamentUC {
	tournaments := make(map[uuid.UUID]*domainObjects.Tournament)
	return &TournamentUC{tournaments, optionUC}
}

func (this *TournamentUC) StartTournament(templateId uuid.UUID) {
	matches := make([]*domainObjects.Match, 0)
	tournament := domainObjects.NewTournament(matches)
	this.tournaments[tournament.GetId()] = tournament
}