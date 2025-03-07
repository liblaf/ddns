package logging

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(level zerolog.Level) {
	log.Logger = log.Level(level).Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
		FormatTimestamp: func(i interface{}) string {
			num, _ := i.(json.Number).Int64()
			timestamp := time.Unix(num, 0)
			return timestamp.Format(time.RFC3339)
		},
	})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
