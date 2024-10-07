package errors

import (
	"log"
	"os"
)

// Used to change the console color for error printing.
const (
	errorColor = "\033[31m"
	colorReset = "\033[0m"
)

// BfErrLogger just encapsulates the default Go logger
// to print error messages as the adopted pattern for
// Butterfly language.
type BfErrLogger struct {
	log       *log.Logger
	errPrefix string
}

// Panic print the args and panics the program.
// Used for fatal errors that must interrupt the
// compiler.
func (bfLog *BfErrLogger) Panicln(args ...any) {
	bfLog.log.Panicln(bfLog.mountArgs(args)...)
}

// Log print the args and continue the program.
// Used for errors that doesn't affect the compilation
// or to accumulate errors before crash.
func (bfLog *BfErrLogger) Println(args ...any) {
	bfLog.log.Println(bfLog.mountArgs(args)...)
}

// Panic print the args and quits the program.
// Used for fatal errors that must interrupt the
// compiler.
func (bfLog *BfErrLogger) Fatalln(args ...any) {
	bfLog.log.Fatalln(bfLog.mountArgs(args)...)
}

// NewBfErrLogger builds and return a new BfErrLogger pointer.
func NewBfErrLogger(errPrefix string) *BfErrLogger {
	var logger = log.New(os.Stderr, errorColor, log.LstdFlags)
	return &BfErrLogger{
		log:       logger,
		errPrefix: errPrefix,
	}
}

// mountArgs build th arguments to the logger adding the message prefix
// defined on the constructor and resetting the console color after
// print the message (otherwise the console would stay red).
func (bfLog *BfErrLogger) mountArgs(args []any) []any {
	var newArgs = []any{bfLog.errPrefix}
	newArgs = append(newArgs, args...)
	newArgs = append(newArgs, colorReset)
	return newArgs
}
