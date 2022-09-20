package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

const (
	LogLevelDisable   = iota
	LogLevelEmergency //logger Panic
	LogLevelAlert     //logger XXX == Panic
	LogLevelCritical  //logger Fatal
	LogLevelError
	LogLevelWarning
	LogLevelNotice //logger XXX == WARING
	LogLevelInfo
	LogLevelDebug
	LogLevelDefault //logger XXX == DEBUG
	LogLevelTrace
)

type LogLevel int

type Logger struct {
	logger   *zerolog.Logger
	logLevel LogLevel
}

var LogLevelToString = [...]string{
	"DISABLE",
	"EMERGENCY",
	"ALERT",
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
	"DEFAULT",
	"TRACE",
}

func (l LogLevel) Name() string { return LogLevelToString[l] }

func GetLogLevelFromString(name string) LogLevel {
	level := LogLevelDisable
	for i, v := range LogLevelToString {
		if v == name {
			level = i
			break
		}
	}
	return LogLevel(level)
}

func New(setLevel LogLevel) *Logger {
	level := zerolog.InfoLevel
	switch setLevel {
	case LogLevelDisable:
		level = zerolog.Disabled
	case LogLevelEmergency:
		fallthrough
	case LogLevelAlert:
		level = zerolog.PanicLevel
	case LogLevelCritical:
		level = zerolog.FatalLevel
	case LogLevelError:
		level = zerolog.ErrorLevel
	case LogLevelWarning:
		fallthrough
	case LogLevelNotice:
		level = zerolog.WarnLevel
	case LogLevelInfo:
		level = zerolog.InfoLevel
	case LogLevelDebug:
		fallthrough
	case LogLevelDefault:
		level = zerolog.DebugLevel
	case LogLevelTrace:
		level = zerolog.TraceLevel
	default:
	}

	//GCP Level
	zerolog.LevelFieldName = "severity"

	zerolog.LevelPanicValue = "EMERGENCY"
	//"ALERT"
	zerolog.LevelFatalValue = "CRITICAL"
	zerolog.LevelErrorValue = "ERROR"
	zerolog.LevelWarnValue = "WARNING"
	//"NOTICE"
	zerolog.LevelInfoValue = "INFO"
	zerolog.LevelDebugValue = "DEBUG"
	//"DEFAULT"

	zerolog.LevelTraceValue = "TRACE"

	zerolog.SetGlobalLevel(level)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{
		logger:   &logger,
		logLevel: setLevel,
	}
}
func (l *Logger) GetHandle() *zerolog.Logger {
	return l.logger
}

// Output duplicates the global logger and sets w as its output.
func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

// With creates a child logger with the field added to its context.
func (l *Logger) With() zerolog.Context {
	return l.logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func (l *Logger) Level(level zerolog.Level) zerolog.Logger {
	return l.logger.Level(level)
}

// Sample returns a logger with the s sampler.
func (l *Logger) Sample(s zerolog.Sampler) zerolog.Logger {
	return l.logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func (l *Logger) Hook(h zerolog.Hook) zerolog.Logger {
	return l.logger.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Err(err error) *zerolog.Event {
	return l.logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Trace() *zerolog.Event {
	return l.logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	l.logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Emergency starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Emergency() *zerolog.Event {
	return l.logger.Panic()
}

// Critical starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Critical() *zerolog.Event {
	return l.logger.Fatal()
}

func (l *Logger) Alert() *zerolog.Event {
	if l.logLevel >= LogLevelAlert {
		return l.logger.Log().Str("severity", "ALERT")
	} else {
		return l.logger.Panic()
	}
}

func (l *Logger) Notice() *zerolog.Event {
	if l.logLevel >= LogLevelNotice {
		return l.logger.Log().Str("severity", "NOTICE")
	} else {
		return l.logger.Info()
	}
}

func (l *Logger) Default() *zerolog.Event {
	if l.logLevel >= LogLevelDefault {
		return l.logger.Log().Str("severity", "DEFAULT")
	} else {
		return l.logger.Trace()
	}
}
