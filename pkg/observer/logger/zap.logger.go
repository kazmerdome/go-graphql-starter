package logger

import "go.uber.org/zap"

type zapLogger struct {
	l *zap.Logger
}

func NewZapLogger() Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &zapLogger{
		l: logger,
	}
}

func (r *zapLogger) Info(msg string)  { r.l.Info(msg) }
func (r *zapLogger) Fatal(msg string) { r.l.Fatal(msg) }
func (r *zapLogger) Warn(msg string)  { r.l.Warn(msg) }
func (r *zapLogger) Error(msg string) { r.l.Error(msg) }
func (r *zapLogger) Panic(msg string) { r.l.Panic(msg) }
