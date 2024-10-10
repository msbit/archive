package main

import (
	"cmp"
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"slices"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s <source> <target>", os.Args[0])
	}

	source := os.Args[1]
	target := os.Args[2]

	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
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

	log.Printf("done")
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
