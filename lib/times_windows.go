package lib

import (
	"errors"
	"io/fs"
	"slices"
	"syscall"
)

func minTimeNs(info fs.FileInfo) (int64, error) {
	stat, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return 0, errors.New("invalid sys")
	}

	return slices.Min([]int64{
		stat.LastAccessTime.Nanoseconds(),
		stat.CreationTime.Nanoseconds(),
		stat.LastWriteTime.Nanoseconds(),
	}), nil
}
