package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Write zerolog.Logger

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	Write = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
func PrintInfo(channel string, message string) {
	Write.Info().Str("Channel", channel).Msg(message)
}
func PrintError(channel string, message string, error error) {
	Write.Error().Str("Channel", channel).Err(error).Msg(message)
}
