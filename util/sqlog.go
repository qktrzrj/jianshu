package util

import (
	"github.com/rs/zerolog"
	"time"
)

type Sqloger struct {
	zerolog.Logger
}

func (l Sqloger) BeforeInFoLog(query string, args ...interface{}) {}

func (l Sqloger) AfterInFoLog(executeTime time.Duration, query string, args ...interface{}) {
	l.Info().Msgf("[SQL] %d %v: %v", executeTime, query, args)
}

func (l Sqloger) Async() bool {
	return false
}

func (l Sqloger) Show() bool {
	return true
}

func NewSQLogger(l zerolog.Logger) Sqloger {
	return Sqloger{l}
}
