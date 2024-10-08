package bfio

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
	errcall "github.com/Hilson-Alex/calldef/src/err_call"
)

const (
	tokenFile          = "doc/Tokens.md"
	goRoot             = "go_compiler"
	goInternalCompiler = "go_compiler/bin/go.exe"
	goFmtInternal      = "go_compiler/bin/gofmt.exe"
	butterflyEmbed     = "butterfly_embed"
	generatedFolder    = "butterfly_embed/generated_code/"
	generatedDocFile   = "bf_____generated.go"
)

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

func HandleFile(name string, handler func(file *os.File)) {
	file, _ := os.Open(name)
	defer errcall.Run(file.Close).OrIgnore()
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

func GenerateGoFile(filename, content string) error {
	var resourceFolder, err = ResourceFolder()
	if err != nil {
		return err
	}
	var tempFile = filepath.Join(resourceFolder, generatedFolder, filename+".go")
	if _, err = os.Stat(tempFile); err == nil {
		return errors.New("Module " + filename + " declared multiple times")
	}
	generated, err := os.Create(tempFile)
	defer errcall.Run(generated.Close).OrIgnore()
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
	var env, args = compilerSetup(resourceFolder)
	return executeExternalEnv(
		env,
		goPath,
		args...,
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

func compilerSetup(baseFolder string) (env, args []string) {
	env = append(
		os.Environ(),
		"GOROOT="+filepath.Join(baseFolder, goRoot),
	)
	args = []string{
		"build",
		"-C=" + filepath.Join(baseFolder, butterflyEmbed),
		"-o=" + outputPath(),
		"-tags=" + RuntimeMode,
		butterflyEmbed,
	}
	return
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
