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
		return fmt.Errorf("unable to get time: %w", err)
	}

	month := t.Format("2006-01")
	target = fmt.Sprintf("%s/%s", target, month)

	err = os.MkdirAll(target, 0750)
	if err != nil {
		return fmt.Errorf("unable to mkdir: %w", err)
	}

	sourceFile := fmt.Sprintf("%s/%s", source, entry.Name())

	r, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("unable to open source: %w", err)
	}
	defer r.Close()

	targetFile := fmt.Sprintf("%s/%s", target, entry.Name())

	if _, err = os.Stat(targetFile); err == nil {
		return fs.ErrExist
	}

	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("unable to stat target: %w", err)
	}

	w, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("unable to create target: %w", err)
	}
	defer w.Close()

	if _, err = io.Copy(w, r); err != nil {
		return fmt.Errorf("unable to copy: %w", err)
	}

	if err := os.Chtimes(targetFile, t, t); err != nil {
		return fmt.Errorf("unable to set times: %w", err)
	}

	return nil
}
