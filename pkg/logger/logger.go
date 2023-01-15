package logger

import "log"

//интерфейс логгера, чтобы не указывать хардкодный Logger. Потом будем писать второй логгер, не придется везде его менять
type Interface interface {
	Info(message string)
	Error(message error)
}

type Logger struct{}

func New() *Logger {
	return new(Logger)
}

func (l *Logger) Info(message string) {
	l.log(message)
}

func (l *Logger) Error(message error) {
	if message == nil {
		return
	}

	l.log(message.Error())
}

func (l *Logger) log(message string) {
	log.Println(message)
}
