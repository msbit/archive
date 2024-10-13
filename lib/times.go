package lib

import (
	"fmt"
	"io/fs"
	"slices"
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

	timestamps, err := timestamps(info)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to get timestamps: %w", err)
	}

	return time.Unix(0, slices.Min(timestamps)), nil
}
