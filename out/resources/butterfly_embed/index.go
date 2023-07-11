package main

import (
	"os"

	"butterfly_embbed/runtime"
)

// silent needed imports
import (
	_ "butterfly_embbed/generated_code"
	_ "butterfly_embbed/system_io"
)

func main() {
	var bootMessage = runtime.BF__MessageCreate("Butterfly Runtime Loader", map[string]interface{}{
		"systemArgs": os.Args,
	})
	runtime.BF__Dispatch("Start", bootMessage)
	runtime.BF__Run()
}
