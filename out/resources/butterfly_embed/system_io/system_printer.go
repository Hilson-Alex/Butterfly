package system_io

import (
	"fmt"

	"butterfly_embed/runtime"
)

func init() {
	const moduleName = "Standard Output"
	runtime.BF__EventSubscribe("Sys Println", moduleName, func(message runtime.BF__MessageContent) {
		switch value := message["text"].(type) {
		case []interface{}:
			fmt.Println(value...)
		default:
			fmt.Println(value)
		}
		if event, ok := message["shares"]; ok {
			runtime.BF__Dispatch(event.(string), runtime.BF__MessageCreate(moduleName, map[string]interface{}{}))
		}
	})
	runtime.BF__EventSubscribe("Sys Printf", moduleName, func(message runtime.BF__MessageContent) {
		var text string
		var args []any
		var ok bool
		if text, ok = message["text"].(string); !ok {
			panic("Printf message is not a string!")
		}
		if args, ok = message["args"].([]interface{}); !ok {
			panic("Printf args are not an array!")
		}
		fmt.Printf(text, args...)
		if event, ok := message["shares"]; ok {
			runtime.BF__Dispatch(event.(string), runtime.BF__MessageCreate(moduleName, map[string]interface{}{}))
		}
	})
}
