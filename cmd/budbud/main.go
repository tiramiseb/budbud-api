package main

import (
	"log"
	"os"

	// "github.com/tiramiseb/budbud-api/internal/delivery/gqlgen"

	"github.com/tiramiseb/budbud-api/internal/authn"
	"github.com/tiramiseb/budbud-api/internal/delivery/gqlgen"
	"github.com/tiramiseb/budbud-api/internal/ownership"
	"github.com/tiramiseb/budbud-api/internal/storage/sqlite"
)

func main() {

	// TODO dependency injection

	s, err := sqlite.New("/tmp/budbud.sqlite")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	authnSrv, err := authn.New(s)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	ownershipSrv, err := ownership.New(s)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	g, err := gqlgen.New(8080, authnSrv, ownershipSrv)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Print("Running...")
	if err := g.Start(); err != nil {
		log.Printf("Failed with error: %v", err)
	}
}
