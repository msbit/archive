package main

import (
	"fmt"
	"io/fs"
	"os"
	"slices"
	"testing"
	"time"
)

func TestCopyDir(t *testing.T) {
	dir := t.TempDir()

	if err := os.MkdirAll(dir+"/source", 0777); err != nil {
		t.Error(err)
	}

	for i := 0; i < 3; i++ {
		if err := os.WriteFile(fmt.Sprintf("%s/source/%02d", dir, i), []byte("foo"), 0777); err != nil {
			t.Error(err)
		}
	}

	for i := 3; i < 6; i++ {
		if err := os.WriteFile(fmt.Sprintf("%s/source/%02d", dir, i), []byte("bar"), 0777); err != nil {
			t.Error(err)
		}
	}

	if err := os.MkdirAll(dir+"/target", 0777); err != nil {
		t.Error(err)
	}

	if err := copyDir(
		dir+"/source",
		dir+"/target",
	); err != nil {
		t.Error(err)
	}

	entries, err := os.ReadDir(dir + "/target/" + time.Now().Format("2006-01"))
	if err != nil {
		t.Error(err)
	}

	names := sliceMap(entries, fs.DirEntry.Name)

	if !slices.Contains(names, "00") {
		t.Error("missing 00")
	}

	if slices.Contains(names, "01") {
		t.Error("contains 01")
	}

	if slices.Contains(names, "02") {
		t.Error("contains 02")
	}

	if !slices.Contains(names, "03") {
		t.Error("missing 03")
	}

	if slices.Contains(names, "04") {
		t.Error("contains 04")
	}

	if slices.Contains(names, "05") {
		t.Error("contains 05")
	}
}

func sliceMap[In any, Out any](s []In, mapper func(In) Out) []Out {
	out := make([]Out, len(s))
	for i, e := range s {
		out[i] = mapper(e)
	}
	return out
}
