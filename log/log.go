package log

import (
	"fmt"
	"github.com/csturiale/go-logger/db"
	"github.com/rs/zerolog"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type logger struct {
	logger        zerolog.Logger
	logExclusions []string
}

var log logger

func Init() error {
	loadFiltersFromCache()
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {

		result := fmt.Sprintf("%s", i)
		for _, e := range log.logExclusions {
			result = strings.ReplaceAll(result, e, "[REDACTED]")
		}
		return result
	}
	buildInfo, _ := debug.ReadBuildInfo()
	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = int(zerolog.InfoLevel) // default to INFO
	}
	w := zerolog.MultiLevelWriter(output)

	log.logger = zerolog.New(w).Level(zerolog.Level(logLevel)).With().Timestamp().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).Logger()

	return nil
}

func loadFiltersFromCache() {
	dbR := db.GetRedisClient()
	if dbR != nil {
		log.logExclusions = dbR.Client.LoadFilters()
	}
}

func saveFilterToCache(msg string) {
	dbR := db.GetRedisClient()
	if dbR != nil {
		dbR.Client.SaveFilter(msg)
	}
}

func Filter(msg string) {
	saveFilterToCache(msg)
	log.logExclusions = append(log.logExclusions, msg)
}

// Debug logs a debug message.
func Debug(msg string) {
	log.logger.Debug().Msg(msg)
}

func Debugf(msg string) {
	log.logger.Debug().Msgf(msg)
}

// Info logs an info message.
func Info(msg string) {
	log.logger.Info().Msg(msg)
}

// Infof Info logs an info message.
func Infof(format string, v ...interface{}) {
	log.logger.Info().Msgf(format, v...)
}

// Warn logs a warning message.
func Warn(msg string) {
	log.logger.Warn().Msg(msg)
}

func Warnf(msg string) {
	log.logger.Warn().Msgf(msg)
}

// Error logs an error message.
func Error(msg string, err error) {
	log.logger.Err(err).Msg(msg)
}

func Errorf(msg string, err error) {
	log.logger.Err(err).Msgf(msg)
}

// Fatal logs a fatal message and exits the program.
func Fatal(msg string, err error) {
	log.logger.Fatal().Err(err).Msg(msg)
}
func Fatalf(msg string, err error) {
	log.logger.Fatal().Err(err).Msgf(msg)
}

// Panic logs a panic message and panics.
func Panic(msg string, err error) {
	log.logger.Panic().Err(err).Msg(msg)
}

func Panicf(msg string, err error) {
	log.logger.Panic().Err(err).Msgf(msg)
}
