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

var generalLogger = bfErrors.NewBfErrLogger("ERROR:")

func main() {
	if bfio.ListTokens {
		listTokens()
		return
	}
	var entries, err = bfio.ReadCompileDir()
	if err != nil {
		generalLogger.Panic(err)
	}

	compileAll(entries)
}

func compileAll(entries []os.DirEntry) {
	var lexerResults = lexAll(entries)
	parseAll(lexerResults)
}

func lexAll(entries []os.DirEntry) [][]*lexer.Token {
	var failedLexers = make([]error, 0, len(entries))
	var parsedTokens = make([][]*lexer.Token, 0, len(entries))
	for _, entry := range entries {
		bfio.HandleFile(filepath.Join(bfio.CompileDir, entry.Name()), func(file *os.File) {
			var tokens, err = lexer.MatchAllTokens(file)
			if err != nil {
				failedLexers = append(failedLexers, err)
			}
			parsedTokens = append(parsedTokens, tokens)
			if bfio.LexerVerbose {
				fmt.Printf("\n%s\n\n", entry.Name())
				for _, token := range tokens {
					fmt.Println(token)
				}
			}
		})
	}
	if len(failedLexers) > 0 {
		generalLogger.Panic(bfErrors.CreateNestedErr("lexical analysis failed", failedLexers...))
	}
	return parsedTokens
}

func parseAll(lexerResults [][]*lexer.Token) {
	var failedParsers = make([]error, 0, len(lexerResults))
	for _, tokens := range lexerResults {
		if !parser.Parse(tokens) {
			failedParsers = append(failedParsers, errors.New("found syntax errorrs on file "+tokens[0].FileName()))
		}
	}
	if len(failedParsers) > 0 {
		generalLogger.Panic(bfErrors.CreateNestedErr("syntax analysis failed", failedParsers...))
	}
	fmt.Println("Syntax analysis completed :)")
}

func listTokens() {
	err := bfio.PrintTokenList()
	if err != nil {
		generalLogger.Panic(err.Error())
	}
}
