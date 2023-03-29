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

const tokenFile = "doc/Tokens.md"

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

func executeExternal(path string) error {
	cmd := exec.Command(path)
	cmd.Stdout = os.Stdout
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
