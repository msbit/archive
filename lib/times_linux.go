package lib

import (
	"errors"
	"io/fs"
	"syscall"
)

func timestamps(info fs.FileInfo) ([]int64, error) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, errors.New("invalid sys")
	}

	return []int64{
		stat.Atim.Nano(),
		stat.Mtim.Nano(),
		stat.Ctim.Nano(),
	}, nil
}
