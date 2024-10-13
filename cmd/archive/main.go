package main

import (
	"log"
	"os"

	"github.com/msbit/archive/lib"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s <source> <target>", os.Args[0])
	}

	if err := lib.CopyDir(
		os.Args[1],
		os.Args[2],
	); err != nil {
		log.Fatal(err)
	}
}
