package lib

import (
	"errors"
	"io/fs"
	"syscall"
)

func timestamps(info fs.FileInfo) ([]int64, error) {
	stat, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return nil, errors.New("invalid sys")
	}

	return []int64{
		stat.LastAccessTime.Nanoseconds(),
		stat.CreationTime.Nanoseconds(),
		stat.LastWriteTime.Nanoseconds(),
	}, nil
}
