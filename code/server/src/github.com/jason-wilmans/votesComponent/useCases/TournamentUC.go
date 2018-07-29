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
	matches			[]*domainObjects.Match
	optionUC		*OptionUC
	random			*rand.Rand
}

func NewTournamentUC(optionUC *OptionUC) *TournamentUC {
	tournaments := make(map[uuid.UUID]*domainObjects.Tournament)
	matches := make([]*domainObjects.Match, 0)
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	return &TournamentUC{tournaments, matches,optionUC, random}
}

func (this *TournamentUC) CreateTournament(templateId uuid.UUID) *domainObjects.Tournament {
	tournament := domainObjects.NewTournament(templateId)
	this.tournaments[tournament.GetId()] = tournament

	options	:= this.optionUC.GetAllForTemplateId(templateId)
	matches := createInitialMatches(tournament.GetId(), options, this.random)
	for _, match := range matches {
		this.matches = append(this.matches, match)
	}

	return tournament
}

func (this *TournamentUC) GetTournament(tournamentId uuid.UUID) *domainObjects.Tournament {
	precond.NotNil(tournamentId, "")

	return this.tournaments[tournamentId]
}

func createInitialMatches(tournamentId uuid.UUID, options []*domainObjects.Option, random *rand.Rand) []*domainObjects.Match {
	nrOfOptions := len(options)
	precond.True(nrOfOptions > 0, "")

	nrOfMatches, highestLevel := calculateNumberOfMatches(nrOfOptions)

	matches := make([]*domainObjects.Match, 0)
	for i := 0; i < nrOfMatches; i++ {
		match := domainObjects.NewMatch(tournamentId, highestLevel, i)
		matches = append(matches, match)
	}

	for _, option := range options {
		matchIndex := getFreeRandomMatch(matches, random)
		matches[matchIndex].Join(option.GetId())
		log.Println("used index ", matchIndex)
	}

	return matches
}

func calculateNumberOfMatches(numberOfOptions int) (int, int) {
	numberOfMatches := 1
	level := 0
	correctedHalf := (numberOfOptions / 2) + (numberOfOptions % 2)
	for numberOfMatches < correctedHalf {
		numberOfMatches = numberOfMatches * 2
		level++
	}
	return numberOfMatches, level
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