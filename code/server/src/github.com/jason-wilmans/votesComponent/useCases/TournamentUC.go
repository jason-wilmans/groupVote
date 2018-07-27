package useCases

import (
	"github.com/satori/go.uuid"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"time"
	"math/rand"
)

type TournamentUC struct {
	tournaments		map[uuid.UUID]*domainObjects.Tournament
	optionUC		*OptionUC
	random			*rand.Rand
}

func NewTournamentUC(optionUC *OptionUC) *TournamentUC {
	tournaments := make(map[uuid.UUID]*domainObjects.Tournament)
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	return &TournamentUC{tournaments, optionUC, random}
}

func (this *TournamentUC) StartTournament(templateId uuid.UUID) *domainObjects.Tournament {
	options	:= this.optionUC.GetAllForTemplateId(templateId)
	matches := createInitialMatches(options, this.random)
	tournament := domainObjects.NewTournament(matches)
	this.tournaments[tournament.GetId()] = tournament
	return tournament
}

func createInitialMatches(options []*domainObjects.Option, random *rand.Rand) []*domainObjects.Match {
	nrOfOptions := len(options)
	nrOfMatches := calculateNumberOfMatches(nrOfOptions)

	matches := make([]*domainObjects.Match, 0)
	for i := 0; i < nrOfMatches; i++ {
		match := domainObjects.NewMatch()
		matches = append(matches, match)
	}

	for _, option := range options {
		placeOptionInRandomMatch(option, matches, random)
	}
	return matches
}

func calculateNumberOfMatches(numberOfOptions int) int {
	numberOfMatches := 1
	for numberOfMatches < numberOfOptions {
		numberOfMatches = numberOfMatches * 2
	}
	return numberOfMatches
}

func placeOptionInRandomMatch(option *domainObjects.Option, matches []*domainObjects.Match, random *rand.Rand) {
	nrOfMatches := len(matches)
	randomIndex := random.Intn(nrOfMatches)
	for !matches[randomIndex].HasRoom() {
		randomIndex = random.Intn(nrOfMatches)
	}
	matches[randomIndex].Join(option.GetId())
}