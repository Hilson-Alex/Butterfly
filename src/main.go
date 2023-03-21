//go:generate goyacc -o ./compiler/parser/butterfly.go ../yacc/butterfly.y

package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Hilson-Alex/Butterfly/src/bfio"
	"github.com/Hilson-Alex/Butterfly/src/compiler/lexer"
	"github.com/Hilson-Alex/Butterfly/src/compiler/parser"
	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

var generalLogger = bfErrors.NewBfErrLogger("ERROR:")

func main() {
	if len(os.Args) < 2 {
		generalLogger.Panic(errors.New("missing the compiling directory"))
	}
	var entries, compileDir, err = bfio.ReadCompileDir()
	if err != nil {
		generalLogger.Panic(err)
	}

	compileAll(entries, compileDir)

	// fmt.Println(dir[0].Name())
	// var reader, err = os.Open(compileDir)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func compileAll(entries []os.DirEntry, dir string) {
	var lexerResults = lexAll(entries, dir)
	for _, tokens := range lexerResults {
		parser.Parse(tokens)
	}
}

func lexAll(entries []os.DirEntry, dir string) [][]*lexer.Token {
	var failedLexers = make([]error, 0)
	var parsedTokens = make([][]*lexer.Token, 0)
	for _, entry := range entries {
		bfio.HandleFile(filepath.Join(dir, entry.Name()), func(file *os.File) {
			var tokens, err = lexer.MatchAllTokens(file)
			if err != nil {
				failedLexers = append(failedLexers, err)
			}
			parsedTokens = append(parsedTokens, tokens)
		})
	}
	if len(failedLexers) > 0 {
		generalLogger.Panic(bfErrors.CreateNestedErr("lexical analysis failed", failedLexers...))
	}
	return parsedTokens
}
