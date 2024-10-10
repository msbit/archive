package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func copy(
	entry fs.DirEntry,
	source string,
	target string,
) error {
	t, err := fsTime(entry)
	if err != nil {
		return err
	}

	month := t.Format("2006-01")
	target = fmt.Sprintf("%s/%s", target, month)

	err = os.MkdirAll(target, 0750)
	if err != nil {
		return err
	}

	sourceFile := fmt.Sprintf("%s/%s", source, entry.Name())

	r, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer r.Close()

	targetFile := fmt.Sprintf("%s/%s", target, entry.Name())

	if _, err = os.Stat(targetFile); err == nil {
		return fs.ErrExist
	}

	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	w, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	if _, err = io.Copy(w, r); err != nil {
		return err
	}

	return os.Chtimes(targetFile, t, t)
}
