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

	stat, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return time.Time{}, errors.New("invalid sys")
	}

	ns := slices.Min([]int64{
		stat.LastAccessTime.Nanoseconds(),
		stat.CreationTime.Nanoseconds(),
		stat.LastWriteTime.Nanoseconds(),
	})

	return time.Unix(0, ns), nil
}
