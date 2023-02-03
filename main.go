package main

import (
	"github.com/kilnfi/near-exporter/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	command := cmd.NewRootCommand()
	if err := command.Execute(); err != nil {
		log.WithError(err).Fatalf("main: execution failed")
	}
}
