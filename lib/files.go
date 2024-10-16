package lib

import (
	"cmp"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"slices"
)

func CopyDir(
	source string,
	target string,
) error {
	entries, err := os.ReadDir(source)
	if err != nil {
		return fmt.Errorf("unable to read dir: %w", err)
	}

	hashes := make(map[string][]fs.DirEntry)
	log.Printf("hashing %d entries (including directories)", len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := fmt.Sprintf("%s/%s", source, entry.Name())
		h, err := hash(name)
		if err != nil {
			return fmt.Errorf("unable to hash file: %w", err)
		}

		hashes[h] = append(hashes[h], entry)
	}

	log.Printf("copying %d unique files", len(hashes))
	for _, entries := range hashes {
		entry := earliest(entries)

		if err := copy(entry, source, target); err != nil {
			log.Printf("not copying %s: %s", entry.Name(), err)
		}
	}

	return nil
}

func hash(name string) (string, error) {
	r, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer r.Close()

	w := md5.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	return string(w.Sum(nil)), nil
}

func earliest(entries []fs.DirEntry) fs.DirEntry {
	return slices.MinFunc(entries, func(a, b fs.DirEntry) int {
		return cmp.Compare(
			mustFsTime(a).UnixNano(),
			mustFsTime(b).UnixNano(),
		)
	})
}

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
