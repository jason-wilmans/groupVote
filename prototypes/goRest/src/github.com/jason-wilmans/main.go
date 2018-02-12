package main

import (
	"github.com/jason-wilmans/infrastructure/routeRegistry"
	"github.com/jason-wilmans/votesComponent"
)

func main() {
	votesComponent.New()
	routeRegistry.StartHosting()
}