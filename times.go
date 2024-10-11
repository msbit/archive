package main

import (
	"fmt"
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

func fsTime(entry fs.DirEntry) (time.Time, error) {
	info, err := entry.Info()
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to get info: %w", err)
	}

	ns, err := minTimeNs(info)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to get minimum timestamp: %w", err)
	}

	return time.Unix(0, ns), nil
}
