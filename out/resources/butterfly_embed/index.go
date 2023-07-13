package main

import (
	"os"

	"butterfly_embed/runtime"
)

// silent needed imports
import (
	_ "butterfly_embed/generated_code"
	_ "butterfly_embed/system_io"
)

func main() {
	var bootMessage = runtime.BF__MessageCreate("Butterfly Runtime Loader", map[string]interface{}{
		"systemArgs": os.Args,
	})
	runtime.BF__Dispatch("Start", bootMessage)
	runtime.BF__Run()
}
