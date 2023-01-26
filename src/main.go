//go:generate goyacc -o ./compiler/test.go ../yacc/test.y

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Hilson-Alex/Butterfly/src/bfio"
	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

func main() {
	var generalLogger = bfErrors.NewBfErrLogger("ERROR:")

	if len(os.Args) < 2 {
		generalLogger.Panic(errors.New("missing the compiling directory"))
	}
	var entries, compileDir, err = bfio.ReadCompileDir()
	if err != nil {
		generalLogger.Panic(err)
	}
	for _, entry := range entries {
		fmt.Println(filepath.Join(compileDir, entry.Name()))
		func() {
			file, _ := os.Open(filepath.Join(compileDir, entry.Name()))
			var scanner = bufio.NewScanner(file)
			scanner.Scan()
			fmt.Println(scanner.Text())
			defer quietClose(file)
		}()
	}
	// fmt.Println(dir[0].Name())
	// var reader, err = os.Open(compileDir)
	// defer func(reader *os.File) {
	// 	_ = reader.Close()
	// }(reader)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func quietClose(file fs.File) {
	_ = file.Close()
}
