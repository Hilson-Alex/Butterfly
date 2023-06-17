//go:build !queue

package runtime

import (
	"fmt"
	"sync"
)

type __bfChannel map[string][]__bfResponse

type __bfCallback func(content BF__MessageContent)

type __bfResponse func(message BF__MessageInfo)

var __eventsRegisteredResponses = make(__bfChannel)

var __bfWg sync.WaitGroup

func BF__EventSubscribe(event, module string, callback __bfCallback) {
	var arr, present = __eventsRegisteredResponses[event]
	if !present {
		arr = make([]__bfResponse, 0)
	}
	__eventsRegisteredResponses[event] = append(arr, __withErrorHandling(event, module, callback))
}

func BF__Dispatch(event string, message BF__MessageInfo) {
	var responses, present = __eventsRegisteredResponses[event]
	if !present {
		fmt.Printf("No responses for event %q\n", event)
		return
	}
	__bfWg.Add(len(responses))
	for _, response := range responses {
		go response(message)
	}
}

func BF__Run() {
	__bfWg.Wait()
}

func __withErrorHandling(event, module string, callback __bfCallback) __bfResponse {
	return func(message BF__MessageInfo) {
		message.EventName = event
		message.ReceiverName = module
		defer func() {
			if r := recover(); r != nil {
				LogError(&message, r)
			}
			__bfWg.Done()
		}()
		callback(message.Content)
	}
}
