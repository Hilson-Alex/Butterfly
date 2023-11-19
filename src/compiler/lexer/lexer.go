package lexer

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/Hilson-Alex/Butterfly/src/bfio"
	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

// Logger for errors when parsing a token.
var logger = bfErrors.NewBfErrLogger("LEXICAL ERROR:")

// MatchAllTokens parse all tokens of a file and return as
// a slice. Returns a slice with all successful tokens.
// error is nil if no unknown token was found (no lexical errors).
// If any unknown token was found, it returns all matched tokens and
// an error.
func MatchAllTokens(file fs.File) ([]*Token, error) {
	var lexer, err = newLexer(file)
	if err != nil {
		return nil, err
	}
	return lexer.parse()
}

// lexer ia a private struct that contains the parsing state
type lexer struct {
	reader       *bufio.Reader
	fileName     string
	line, column int
	currentLine  string
}

// creates a new lexer from a file returns an error if
// it could not read the file stats
func newLexer(file fs.File) (*lexer, error) {
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &lexer{
		reader:   bufio.NewReader(file),
		fileName: filepath.Join(bfio.CompileDir, stats.Name()),
		line:     1,
		column:   1,
	}, nil
}

func safeRead(reader *bufio.Reader) (string, bool) {
	var EOF = false
	var text = bfErrors.Resolve(func() (string, error) { return reader.ReadString('\n') }).
		AndHandle(func(err error) {
			if err != io.EOF {
				logger.Panic(err)
			}
			EOF = true
		})
	return text, EOF
}

// parse gets all valid tokens from the lexer reader.
// if an unknown token is found, it logs the occurrence
// and continue the parsing.
func (lexer *lexer) parse() ([]*Token, error) {
	var reader = lexer.reader
	var tokens = make([]*Token, 0)
	var hasError = false
	var err error
	var EOF = false
	for lexer.currentLine, EOF = safeRead(reader); ; lexer.currentLine, EOF = safeRead(reader) {
		for lexer.currentLine != "" {
			if lexer.skippedText() {
				continue
			}
			if token := lexer.parseNextToken(); token != nil {
				tokens = append(tokens, token)
				continue
			}
			hasError = true
			lexer.logError()
		}
		lexer.column = 1
		if EOF {
			break
		}
	}
	err = nil
	if hasError {
		err = errors.New("found unknown tokens for file " + lexer.fileName)
	}
	return tokens, err
}

// parseNextToken parses the next token of the lexer current line
func (lexer *lexer) parseNextToken() *Token {
	var text, line, column = lexer.currentLine, lexer.line, lexer.column
	for _, token := range matchers {
		if str := token.matcher.FindString(text); str != "" {
			lexer.currentLine = text[len(str):]
			lexer.column += len(str)
			return &Token{
				tokenType: token.code,
				value:     str,
				line:      line,
				column:    column,
				fileName:  lexer.fileName,
			}
		}
	}
	return nil
}

// skippedText skips and return a flag if the following characters
// can be ignored
func (lexer *lexer) skippedText() bool {
	var text = lexer.currentLine
	if text == "\n" || text == "\r\n" || lineComment.MatchString(text) {
		lexer.currentLine = ""
		lexer.line++
		return true
	}
	if text[0] == ' ' || text[0] == '\t' {
		lexer.currentLine = text[1:]
		lexer.column++
		return true
	}
	return false
}

// logError takes the next unknown token and logs a lexical
// error.
func (lexer *lexer) logError() {
	var text, line, column = lexer.currentLine, lexer.line, lexer.column
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
}
