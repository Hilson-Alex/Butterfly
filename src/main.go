//go:generate goyacc -o ./compiler/parser/butterfly.go ../yacc/butterfly.y

package main

import (
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
	var parseResults = parseAll(lexerResults)
	generateAll(parseResults)
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

func parseAll(lexerResults [][]*lexer.Token) []parser.ParserResult {
	var failed = false
	var results = make([]parser.ParserResult, 0, len(lexerResults))
	for _, tokens := range lexerResults {
		if len(tokens) == 0 {
			continue
		}
		var parserResult = parser.Parse(tokens)
		if !parserResult.Success {
			failed = true
			continue
		}
		results = append(results, parserResult)
	}
	if failed {
		generalLogger.Panic("syntax analysis failed")
	}
	return results
}

func generateAll(parserResults []parser.ParserResult) {
	var failedFiles = make([]error, 0, len(parserResults))
	bfio.CleanGeneratedFiles()
	for _, result := range parserResults {
		if err := bfio.GenerateGoFile(result.ModuleName, result.Result); err != nil {
			failedFiles = append(failedFiles, err)
		}
	}
	if len(failedFiles) > 0 {
		generalLogger.Panic(bfErrors.CreateNestedErr("error generating files", failedFiles...))
	}
	if err := bfio.GoCompile(); err != nil {
		generalLogger.Panic(bfErrors.CreateNestedErr("error generating executable", err))
	}
	if !bfio.KeepGenFiles {
		bfio.CleanGeneratedFiles()
		return
	}
	_ = bfio.GoFmtGenerated()
}

func listTokens() {
	err := bfio.PrintTokenList()
	if err != nil {
		generalLogger.Panic(err.Error())
	}
}
