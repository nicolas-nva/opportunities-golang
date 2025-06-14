package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer	io.Writer
}

func NewLogger(prfx string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prfx, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer,"DEBUG: ", logger.Flags()),
		info:	 log.New(writer,"INFO: ", logger.Flags()),
		warning: log.New(writer,"WARNING: ", logger.Flags()),
		err: 	 log.New(writer,"ERROR: ", logger.Flags()),
		writer: writer,
	}
}

// create non-formated logs
func (l *Logger) Debug(v ... interface{}){
	l.debug.Println(v ...)
}

func (l *Logger) Info(v ... interface{}){
	l.info.Println(v ...)
}
func (l *Logger) Warn(v ... interface{}){
	l.warning.Println(v ...)
}
func (l *Logger) Error(v ... interface{}){
	l.err.Println(v ...)
}

//create format enabled logs
func (l *Logger) DebugF(format string, v... interface{}){
	l.debug.Printf(format, v...)
}

func (l *Logger) InfoF(format string, v... interface{}){
	l.info.Printf(format, v...)
}

func (l *Logger) WarnF(format string, v... interface{}){
	l.warning.Printf(format, v...)
}

func (l *Logger) ErrorF(format string, v... interface{}){
	l.err.Printf(format, v...)
}