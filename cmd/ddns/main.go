package main

import (
	"github.com/liblaf/ddns/pkg/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	cmd := cmd.Root()
	if err := cmd.Execute(); err != nil {
		log.Fatal().Msgf("%+v", err)
	}
}
