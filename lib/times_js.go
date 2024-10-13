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
		(&syscall.Timespec{
			stat.Atime,
			stat.AtimeNsec,
		}).Nano(),
		(&syscall.Timespec{
			stat.Mtime,
			stat.MtimeNsec,
		}).Nano(),
		(&syscall.Timespec{
			stat.Ctime,
			stat.CtimeNsec,
		}).Nano(),
	}, nil
}
