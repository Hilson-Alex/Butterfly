//go:generate goyacc -o ./compiler/parser/butterfly.go ../yacc/butterfly.y

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Hilson-Alex/Butterfly/src/bfio"
	"github.com/Hilson-Alex/Butterfly/src/compiler/lexer"
	"github.com/Hilson-Alex/Butterfly/src/compiler/parser"
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
		func(name string) {
			file, _ := os.Open(filepath.Join(compileDir, name))
			defer bfio.QuietClose(file)
			var tokens = lexer.MatchAllTokens(file)
			fmt.Println(name)
			var inter = parser.Parser[*lexer.Token]{TokBuffer: tokens}
			var token = new(*lexer.Token)
			for code := inter.Lex(token); code != 0; code = inter.Lex(token) {
				fmt.Println(*token)
			}
			fmt.Println()
		}(entry.Name())
	}
	// fmt.Println(dir[0].Name())
	// var reader, err = os.Open(compileDir)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
