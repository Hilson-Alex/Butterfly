package bfio

import (
	"flag"
)

var OutputPath string

var ListTokens bool

var LexerVerbose bool

var CompileDir string

func init() {
	flag.StringVar(&OutputPath, "out", "", "Gives a custom path to the generated exe")
	flag.BoolVar(&ListTokens, "listTokens", false, "List all language tokens and their use. Useful to understand syntax errors")
	flag.BoolVar(&LexerVerbose, "lexerVerbose", false, "Prints the parsed Tokens while compiling")
	flag.Parse()
	CompileDir = flag.Arg(0)
}
