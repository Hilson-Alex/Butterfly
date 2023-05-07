package bfio

import (
	"flag"
	"fmt"
	"os"
)

var OutputPath string

var ListTokens bool

var LexerVerbose bool

var KeepGenFiles bool

var CompileDir string

func init() {
	flag.StringVar(&OutputPath, "out", "", "Gives a custom path to the generated exe")
	flag.BoolVar(&ListTokens, "listTokens", false, "List all language tokens and their use. Useful to understand syntax errors")
	flag.BoolVar(&LexerVerbose, "lexerVerbose", false, "Prints the parsed Tokens while compiling")
	flag.BoolVar(&KeepGenFiles, "keepGeneratedFiles", false, "Don't delete the generated files after compiling")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: butterfly [flags] [compiling directory]\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Flags for %s:\n", os.Args[0])
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "  -help\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "    \tPrints this message\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	CompileDir = flag.Arg(0)
}
