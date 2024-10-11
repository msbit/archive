package main

import (
	"io/fs"
	"time"
)

func mustFsTime(entry fs.DirEntry) time.Time {
	t, err := fsTime(entry)
	if err != nil {
		panic(err)
	}

	return t
}
