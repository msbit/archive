package main

import (
	"errors"
	"io/fs"
	"slices"
	"syscall"
	"time"
)

func fsTime(entry fs.DirEntry) (time.Time, error) {
	info, err := entry.Info()
	if err != nil {
		return time.Time{}, err
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return time.Time{}, errors.New("invalid sys")
	}

	ns := slices.Min([]int64{
		stat.Atimespec.Nano(),
		stat.Mtimespec.Nano(),
		stat.Ctimespec.Nano(),
		stat.Birthtimespec.Nano(),
	})

	return time.Unix(0, ns), nil
}
