package logger

type Logger interface {
	Info(msg string)
	Fatal(msg string)
	Warn(msg string)
	Error(msg string)
	Panic(msg string)
}
