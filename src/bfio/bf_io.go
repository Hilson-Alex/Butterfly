package bfio

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"

	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
	bfEmbed "github.com/Hilson-Alex/butterfly_embed"
	"github.com/Hilson-Alex/goembed/compiler"
	"github.com/Hilson-Alex/goembed/props"
)

const tokenFile = "doc/Tokens.md"

func ReadCompileDir() (entries []os.DirEntry, err error) {
	if CompileDir == "" {
		err = errors.New("missing the compiling directory. Please inform it like: `butterfly ./example`")
		return
	}
	entries, err = os.ReadDir(CompileDir)
	if err == nil {
		entries = filterEntries(entries)
		return
	}
	if errors.Is(err, syscall.ENOTDIR) {
		var errMsg = "could not open the specified path as an folder. Are you sure it is a valid directory?"
		err = bfErrors.CreateNestedErr(errMsg, err)
		return
	}
	err = bfErrors.CreateNestedErr("can't open directory", err)
	return
}

func HandleFile(name string, handler func(file *os.File)) {
	file, _ := os.Open(name)
	defer func() { _ = file.Close() }()
	handler(file)
}

func ResourceFolder() (string, error) {
	var exePath, err = os.Executable()
	if err != nil {
		return "", err
	}
	exePath = filepath.Dir(exePath)
	return filepath.Join(exePath, "resources"), nil
}

func BuildRuntime() (*bfEmbed.BFRuntime, error) {
	var resourceFolder, err = ResourceFolder()
	if err != nil {
		return nil, err
	}
	return bfEmbed.CreateRuntime(resourceFolder)
}

func GoCompile(bfRuntime *bfEmbed.BFRuntime) error {
	var resourceFolder, err = ResourceFolder()
	if err != nil {
		return err
	}
	compiler.CacheRoot = resourceFolder
	return compiler.GoBuild(
		bfRuntime.GoModule(),
		compiler.BuildOptions(
			compiler.WithArgs(
				props.From(map[string]string{
					"-C":    bfRuntime.CompileDir(),
					"-o":    outputPath(),
					"-tags": RuntimeMode,
				}),
			),
		),
	)
}

func GoFmtGenerated(bfRuntime *bfEmbed.BFRuntime) error {
	return compiler.GoFmt(
		bfRuntime.GenerateDir(),
		compiler.BuildOptions(
			compiler.WithArgs(
				props.From(map[string]string{
					"-s": "",
					"-w": "",
				}),
			),
		),
	)
}

func filterEntries(entries []os.DirEntry) []os.DirEntry {
	var filteredEntries = make([]os.DirEntry, 0)
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".bf" {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return filteredEntries
}

func outputPath() string {
	var absPath string
	var exeName = filepath.Base(CompileDir) + ".exe"
	if OutputPath == "" {
		absPath, _ = filepath.Abs(CompileDir)
		return filepath.Join(absPath, exeName)
	}
	absPath, _ = filepath.Abs(OutputPath)
	if filepath.Ext(absPath) != ".exe" {
		return filepath.Join(absPath, exeName)
	}
	exeName = filepath.Base(absPath)
	absPath, _ = filepath.Abs(filepath.Dir(absPath))
	return filepath.Join(absPath, exeName)
}
