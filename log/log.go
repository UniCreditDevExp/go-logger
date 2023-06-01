package log

import (
	"fmt"
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

var Log logger

func Init() error {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {

		result := fmt.Sprintf("%s", i)
		for _, e := range Log.logExclusions {
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

	Log.logger = zerolog.New(w).Level(zerolog.Level(logLevel)).With().Timestamp().Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).Logger()

	return nil
}

func (l *logger) Filter(s string) {
	l.logExclusions = append(l.logExclusions, s)
}

// Debug logs a debug message.
func Debug(msg string) {
	Log.logger.Debug().Msg(msg)
}

// Info logs an info message.
func Info(msg string) {
	Log.logger.Info().Msg(msg)
}

// Infof Info logs an info message.
func Infof(format string, v ...interface{}) {
	Log.logger.Info().Msgf(format, v...)
}

// Warn logs a warning message.
func Warn(msg string) {
	Log.logger.Warn().Msg(msg)
}

// Error logs an error message.
func Error(msg string, err error) {
	Log.logger.Err(err).Msg(msg)
}

// Fatal logs a fatal message and exits the program.
func Fatal(msg string, err error) {
	Log.logger.Fatal().Err(err).Msg(msg)
}

// Panic logs a panic message and panics.
func Panic(msg string, err error) {
	Log.logger.Panic().Err(err).Msg(msg)
}
