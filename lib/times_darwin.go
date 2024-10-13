package lib

import (
	"errors"
	"io/fs"
	"slices"
	"syscall"
)

func minTimeNs(info fs.FileInfo) (int64, error) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, errors.New("invalid sys")
	}

	return slices.Min([]int64{
		stat.Atimespec.Nano(),
		stat.Mtimespec.Nano(),
		stat.Ctimespec.Nano(),
		stat.Birthtimespec.Nano(),
	}), nil
}
