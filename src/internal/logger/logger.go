package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"movie-journal-api/config"
	"os"
	"strconv"
	"sync"
	"time"
)

var logger zerolog.Logger
var once sync.Once

func Initialize() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.InfoLevel) // default to INFO
		}

		var output io.Writer = zerolog.MultiLevelWriter(os.Stderr)
		currentDate := time.Now().Format(config.DateFormat)
		logFileName := fmt.Sprintf("storage/logs/%s.log", currentDate)

		if config.FileLog == "true" {
			fileLogger := &lumberjack.Logger{
				Filename:   logFileName,
				MaxSize:    5, //
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}

		logger = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Logger()
	})
}
func PrintInfo(channel string, message string) {
	logger.Info().Str("Channel", channel).Msg(message)
}
func PrintError(channel string, message string, trace string, err error) {
	log := logger.Error().Str("Channel", channel)
	if trace != "" {
		log = log.Str("Trace", trace)
	}
	if err != nil {
		log = log.Err(err)
	}

	log.Msg(message)
}
