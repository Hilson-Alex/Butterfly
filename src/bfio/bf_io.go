package bfio

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

type callback func(file *os.File)

const (
	tokenFile          = "doc/Tokens.md"
	goRoot             = "go_compiler"
	goInternalCompiler = "go_compiler/bin/go.exe"
	goFmtInternal      = "go_compiler/bin/gofmt.exe"
	butterflyEmbbed    = "butterfly_embbed"
	generatedFolder    = "butterfly_embbed/generated_code/"
	generatedDocFile   = "bf_____generated.go"
)

func ReadCompileDir() (entries []os.DirEntry, err error) {
	if CompileDir == "" {
		err = errors.New("missing the compiling directory. Please inform it like: `butterfly ./example`")
		return
	}
	entries, err = os.ReadDir(CompileDir)
	if err != nil {
		if errors.Is(err, syscall.ENOTDIR) {
			var errMsg = "could not open the specified path as an folder. Are you sure it is a directory?"
			err = bfErrors.CreateNestedErr(errMsg, err)
			return
		}
		err = bfErrors.CreateNestedErr("can't open directory", err)
		return
	}
	entries = filterEntries(entries)
	return
}

func PrintTokenList() error {
	var resourceFolder, err = ResourceFolder()
	if err != nil {
		return err
	}
	HandleFile(filepath.Join(resourceFolder, tokenFile), func(file *os.File) {
		var scanner = bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(strings.ReplaceAll(scanner.Text(), "**", ""))
		}
		err = scanner.Err()
	})
	return err
}

func HandleFile(name string, handler callback) {
	file, _ := os.Open(name)
	defer QuietClose(file)
	handler(file)
}

func QuietClose(file fs.File) {
	_ = file.Close()
}

func ResourceFolder() (string, error) {
	var exePath, err = os.Executable()
	if err != nil {
		return "", err
	}
	exePath = filepath.Dir(exePath)
	return filepath.Join(exePath, "resources"), nil
}

func GenerateGoFile(filename, content string) error {
	var resourceFolder, err = ResourceFolder()
	if err != nil {
		return err
	}
	generated, err := os.Create(filepath.Join(resourceFolder, generatedFolder, filename+".go"))
	defer QuietClose(generated)
	if err != nil {
		return err
	}
	_, err = generated.WriteString(content)
	return err
}

func GoCompile() error {
	var resourceFolder, err = ResourceFolder()
	if err != nil {
		return err
	}
	var goPath = filepath.Join(resourceFolder, goInternalCompiler)
	var env = append(
		os.Environ(),
		"GOROOT="+filepath.Join(resourceFolder, goRoot),
	)
	return executeExternalEnv(
		env,
		goPath,
		"build",
		"-C="+filepath.Join(resourceFolder, butterflyEmbbed),
		"-o="+outputPath(),
		butterflyEmbbed,
	)
}

func GoFmtGenerated() error {
	var resourceFolder, err = ResourceFolder()
	if err != nil {
		return err
	}
	var generated = filepath.Join(resourceFolder, generatedFolder)
	var goFmt = filepath.Join(resourceFolder, goFmtInternal)
	return executeExternal(goFmt, "-s", "-w", generated)
}

func CleanGeneratedFiles() {
	var resourceFolder, _ = ResourceFolder()
	var entries, _ = os.ReadDir(filepath.Join(resourceFolder, generatedFolder))
	for _, file := range entries {
		if file.Name() != generatedDocFile {
			_ = os.Remove(filepath.Join(resourceFolder, generatedFolder, file.Name()))
		}
	}
}

func executeExternal(path string, arg ...string) error {
	return executeExternalEnv(os.Environ(), path, arg...)
}

func executeExternalEnv(env []string, path string, arg ...string) error {
	cmd := exec.Command(path, arg...)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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
	return absPath
}
