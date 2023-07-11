//go:build queue

package runtime

import (
	"fmt"
	"sync"
)

type __bfChannel map[string][]__bfResponse

type __bfCallback func(content BF__MessageContent)

type __bfResponse func(message BF__MessageInfo)

type __syncArr struct {
	mutex sync.Mutex
	arr   []BF__MessageInfo
}

var __eventsRegisteredResponses = make(__bfChannel)

var __eventQueue = __syncArr{arr: make([]BF__MessageInfo, 0)}

var __bfWg sync.WaitGroup

func BF__EventSubscribe(event, module string, callback __bfCallback) {
	var arr, present = __eventsRegisteredResponses[event]
	if !present {
		arr = make([]__bfResponse, 0)
	}
	__eventsRegisteredResponses[event] = append(arr, __withErrorHandling(event, module, callback))
}

func BF__Dispatch(event string, message BF__MessageInfo) {
	message.EventName = event
	__eventQueue.mutex.Lock()
	__eventQueue.arr = append(__eventQueue.arr, message)
	__eventQueue.mutex.Unlock()
}

func BF__Run() {
	for len(__eventQueue.arr) != 0 {
		__runEvent()
	}
}

func __runEvent() {
	var message = __eventQueue.arr[0]
	__eventQueue.arr = __eventQueue.arr[1:]
	var event = message.EventName
	var responses, present = __eventsRegisteredResponses[event]
	if !present {
		fmt.Printf("No responses for event %q\n", event)
		return
	}
	__bfWg.Add(len(responses))
	for _, response := range responses {
		go response(message)
	}
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
