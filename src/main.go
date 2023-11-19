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

// aliases
type lexerList = [][]*lexer.Token
type parserList = []parser.ParserResult

func main() {
	if bfio.ListTokens {
		bfErrors.Run(bfio.PrintTokenList).OrPanicOn(generalLogger)
		return
	}
	var entries = bfErrors.Resolve(bfio.ReadCompileDir).OrPanicOn(generalLogger)
	compileAll(entries)
}

func compileAll(entries []os.DirEntry) {

	var lexerResults = bfErrors.
		Resolve(func() (lexerList, error) { return lexAll(entries) }).
		OrPanicOn(generalLogger)

	var parseResults = bfErrors.
		Resolve(func() (parserList, error) { return parseAll(lexerResults) }).
		OrPanicOn(generalLogger)

	bfErrors.
		Run(func() error { return generateAll(parseResults) }).
		OrPanicOn(generalLogger)
}

func lexAll(entries []os.DirEntry) (lexerList, error) {
	var failedLexers = make([]error, 0, len(entries))
	var parsedTokens = make(lexerList, 0, len(entries))
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
		return nil, bfErrors.CreateNestedErr("lexical analysis failed", failedLexers...)
	}
	return parsedTokens, nil
}

func parseAll(lexerResults lexerList) ([]parser.ParserResult, error) {
	var failedParsers = make([]error, 0, len(lexerResults))
	var results = make([]parser.ParserResult, 0, len(lexerResults))
	for _, tokens := range lexerResults {
		if len(tokens) == 0 {
			continue
		}
		var parserResult = parser.Parse(tokens)
		if !parserResult.Success {
			failedParsers = append(failedParsers, errors.New("found syntax errorrs on file "+tokens[0].FileName()))
		}
		results = append(results, parserResult)
	}
	if len(failedParsers) > 0 {
		return nil, bfErrors.CreateNestedErr("syntax analysis failed", failedParsers...)
	}
	return results, nil
}

func generateAll(parserResults []parser.ParserResult) error {
	var failedFiles = make([]error, 0, len(parserResults))
	bfio.CleanGeneratedFiles()
	for _, result := range parserResults {
		bfErrors.
			Run(func() error { return bfio.GenerateGoFile(result.ModuleName, result.Result) }).
			AndHandle(func(err error) { failedFiles = append(failedFiles, err) })
	}

	if len(failedFiles) > 0 {
		return bfErrors.CreateNestedErr("error generating files", failedFiles...)
	}

	if err := bfio.GoCompile(); err != nil {
		return bfErrors.CreateNestedErr("error generating executable", err)
	}
	if !bfio.KeepGenFiles {
		bfio.CleanGeneratedFiles()
		return nil
	}
	_ = bfio.GoFmtGenerated()
	return nil
}
