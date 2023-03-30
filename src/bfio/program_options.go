package bfio

import (
	"flag"
	"fmt"
	"os"
)

var OutputPath string

var ListTokens bool

var LexerVerbose bool

var CompileDir string

func init() {
	flag.StringVar(&OutputPath, "out", "", "Gives a custom path to the generated exe")
	flag.BoolVar(&ListTokens, "listTokens", false, "List all language tokens and their use. Useful to understand syntax errors")
	flag.BoolVar(&LexerVerbose, "lexerVerbose", false, "Prints the parsed Tokens while compiling")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: butterfly [flags] [compiling directory]\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Flags for %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  -help\n")
		fmt.Fprintf(flag.CommandLine.Output(), "    \tPrints this message\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	CompileDir = flag.Arg(0)
}
