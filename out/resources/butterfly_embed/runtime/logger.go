package runtime

import (
	"log"
	"os"
)

const (
	errorColor  = "\033[31m"
	errorPrefix = "BUTTERFLY RUNTIME ERROR"
	colorReset  = "\033[0m"
)

var logger = log.New(os.Stderr, errorColor, log.LstdFlags)

func LogError(message *BF__MessageInfo, description any) {
	logger.Printf("%s: %v\n\tevent: %s"+
		"\n\tmessage sender: %s"+
		"\n\tmessage receiver: %s%s\n",
		errorPrefix, description,
		message.EventName, message.SenderName, message.ReceiverName,
		colorReset,
	)
}
