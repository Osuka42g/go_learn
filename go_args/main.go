package main

import (
	"fmt"
	"io"
	"os"
)

type readFile struct{}

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Insufficient arguments")
		os.Exit(1)
	}

	file, err := os.Open(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rf := readFile{}
	io.Copy(rf, file)
}

func (readFile) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return 0, nil
}
