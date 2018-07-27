package useCases

import (
	"github.com/satori/go.uuid"
	"github.com/jason-wilmans/votesComponent/domainObjects"
	"time"
	"math/rand"
	"github.com/muroc/prego"
	"log"
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

func (this *TournamentUC) CreateTournament(templateId uuid.UUID) *domainObjects.Tournament {
	options	:= this.optionUC.GetAllForTemplateId(templateId)
	matches := createInitialMatches(options, this.random)
	tournament := domainObjects.NewTournament(matches)
	this.tournaments[tournament.GetId()] = tournament
	return tournament
}

func (this *TournamentUC) GetTournament(tournamentId uuid.UUID) *domainObjects.Tournament {
	precond.NotNil(tournamentId, "")

	return this.tournaments[tournamentId]
}

func createInitialMatches(options []*domainObjects.Option, random *rand.Rand) []*domainObjects.Match {
	nrOfOptions := len(options)
	precond.True(nrOfOptions > 0, "")

	nrOfMatches := calculateNumberOfMatches(nrOfOptions)

	matches := make([]*domainObjects.Match, 0)
	for i := 0; i < nrOfMatches; i++ {
		match := domainObjects.NewMatch()
		matches = append(matches, match)
	}

	for _, option := range options {
		matchIndex := getFreeRandomMatch(matches, random)
		matches[matchIndex].Join(option.GetId())
		log.Println("used index ", matchIndex)
	}

	return matches
}

func calculateNumberOfMatches(numberOfOptions int) int {
	numberOfMatches := 1
	correctedHalf := (numberOfOptions / 2) + (numberOfOptions % 2)
	for numberOfMatches < correctedHalf {
		numberOfMatches = numberOfMatches * 2
	}
	return numberOfMatches
}

func getFreeRandomMatch(matches []*domainObjects.Match, random *rand.Rand) int {
	nrOfMatches := len(matches)

	randomIndex := getRandomIndex(nrOfMatches, random)
	for !matches[randomIndex].HasRoom() {
		if nrOfMatches == 1 {
			panic("Something went horribly wrong :(")
		}
		randomIndex = getRandomIndex(nrOfMatches, random)
	}
	return randomIndex
}

func getRandomIndex(nrOfMatches int, random *rand.Rand) int {
	if nrOfMatches > 1 {
		return random.Intn(nrOfMatches)
	} else {
		return 0
	}
}