package logger

import (
	"log"
	"os"
)

type (
	standardLogger struct {
		l *log.Logger
	}
)

func NewStandardLogger() Logger {
	logger := log.New(os.Stdout, "", 0)
	return &standardLogger{
		l: logger,
	}
}

func (r *standardLogger) Info(msg string)  { r.l.Printf(msg) }
func (r *standardLogger) Fatal(msg string) { r.l.Fatal(msg) }
func (r *standardLogger) Warn(msg string)  { r.l.Printf(msg) }
func (r *standardLogger) Error(msg string) { r.l.Fatal(msg) }
func (r *standardLogger) Panic(msg string) { r.l.Panic(msg) }
