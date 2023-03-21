package bfio

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

type callback = func(file *os.File)

func ReadCompileDir() (entries []os.DirEntry, root string, err error) {
	if len(os.Args) < 2 {
		err = errors.New("missing the compiling directory")
		return
	}
	root = os.Args[1]
	entries, err = os.ReadDir(root)
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

func QuietClose(file fs.File) {
	_ = file.Close()
}

func HandleFile(name string, handler callback) {
	file, _ := os.Open(name)
	defer QuietClose(file)
	handler(file)
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
