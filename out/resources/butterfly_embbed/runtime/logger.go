package runtime

import (
	"log"
)

const (
	errorColor  = "\033[31m"
	errorPrefix = "BUTTERFLY RUNTIME ERROR"
	colorReset  = "\033[0m"
)

var logger = log.Default()

func LogError(message *BF__MessageInfo, description any) {
	logger.Printf("%s%s: %v\n\tevent: %s"+
		"\n\tmessage sender: %s"+
		"\n\tmessage receiver: %s%s\n",
		errorColor, errorPrefix, description,
		message.EventName, message.SenderName, message.ReceiverName,
		colorReset,
	)
}
