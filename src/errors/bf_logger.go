package errors

import (
	"log"
	"os"
)

const (
	errorColor = "\033[31m"
	colorReset = "\033[0m"
)

type BfErrLogger struct {
	log       *log.Logger
	errPrefix string
}

func (bfLog *BfErrLogger) Panic(args ...any) {
	bfLog.log.Fatalln(bfLog.mountArgs(args)...)
}

func (bfLog *BfErrLogger) Log(args ...any) {
	bfLog.log.Println(bfLog.mountArgs(args)...)
}

func NewBfErrLogger(errPrefix string) *BfErrLogger {
	var logger = log.New(os.Stderr, errorColor, log.LstdFlags)
	return &BfErrLogger{
		log:       logger,
		errPrefix: errPrefix,
	}
}

func (bfLog *BfErrLogger) mountArgs(args []any) []any {
	var newArgs = []any{bfLog.errPrefix}
	newArgs = append(newArgs, args...)
	newArgs = append(newArgs, colorReset)
	return newArgs
}
