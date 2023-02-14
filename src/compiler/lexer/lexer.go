package lexer

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

var logger = bfErrors.NewBfErrLogger("LEXICAL ERROR")

func MatchAllTokens(file fs.File) []*Token {
	return newLexer(file).parse()
}

type lexer struct {
	reader       *bufio.Reader
	fileName     string
	line, column int
	currentLine  string
}

func newLexer(file fs.File) *lexer {
	stats, _ := file.Stat()
	return &lexer{
		reader:   bufio.NewReader(file),
		fileName: filepath.Join(os.Args[1], stats.Name()),
		line:     1,
		column:   1,
	}
}

func (lexer *lexer) parse() []*Token {
	var reader = lexer.reader
	var tokens = make([]*Token, 0)
	var err error
	for lexer.currentLine, err = reader.ReadString('\n'); ; lexer.currentLine, err = reader.ReadString('\n') {
		for lexer.currentLine != "" {
			if token := lexer.parseNextToken(); token != nil {
				tokens = append(tokens, token)
			}
		}
		lexer.column = 1
		if err == io.EOF {
			break
		}
	}
	return tokens
}

func (lexer *lexer) parseNextToken() *Token {
	var text, line, column = lexer.currentLine, lexer.line, lexer.column
	if text[0] == '\n' || text == "\r\n" || lineComment.MatchString(text) {
		lexer.currentLine = ""
		lexer.line++
		return nil
	}
	if text[0] == ' ' || text[0] == '\t' {
		lexer.currentLine = text[1:]
		lexer.column++
		return nil
	}
	for _, token := range matchers {
		if str := token.matcher.FindString(text); str != "" {
			str = strings.TrimSpace(str)
			lexer.currentLine = text[len(str):]
			lexer.column += len(str)
			return &Token{
				tokenType: token.code,
				value:     str,
				line:      line,
				column:    column,
			}
		}
	}
	var invalidToken = getUnknownToken(text)
	var offset = len(invalidToken)
	logger.Log(&bfErrors.TokenError{
		Line:     line,
		Column:   column,
		Offset:   offset,
		Value:    invalidToken,
		FileName: lexer.fileName,
	})
	if offset == len(text) {
		lexer.currentLine = ""
	} else {
		lexer.currentLine = text[offset:]
	}
	return nil
}
