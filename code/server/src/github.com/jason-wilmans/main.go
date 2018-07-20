package main

import (
	"github.com/jason-wilmans/infrastructure/routeRegistry"
	"github.com/jason-wilmans/webAPI"
	"github.com/jason-wilmans/votesComponent"
)

var vote = votesComponent.New()
var _ = webAPI.NewVoteController(vote)

func main() {
	routeRegistry.StartHosting()
}