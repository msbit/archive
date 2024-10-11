package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s <source> <target>", os.Args[0])
	}

	if err := copyDir(
		os.Args[1],
		os.Args[2],
	); err != nil {
		log.Fatal(err)
	}
}
