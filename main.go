package main

import (
	"flag"
	"fmt"
	"os"
)

func help() {
	fmt.Println("Usage:")
	fmt.Println("  go run ./testing/filter -dir {string} -sequence {number} -parallel {number}")
	fmt.Println("")
	fmt.Println("  dir:      target root directory")
	fmt.Println("  sequence: a number to specify running machine")
	fmt.Println("  parallel: max parallelism")
	fmt.Println("")
}

func parseFlags() (root string, sequence, parallel int) {
	flag.StringVar(&root, "dir", "", "")
	flag.IntVar(&sequence, "sequence", -1, "")
	flag.IntVar(&parallel, "parallel", -1, "")

	flag.Parse()

	if root == "" || sequence < 0 || parallel < 1 {
		help()

		panic("invalid params")
	}

	return root, sequence, parallel
}

func main() {
	root, sequence, parallel := parseFlags()

	dirs, err := readdir(root)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	chunks, err := chunk(dirs, sequence, parallel)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(chunks)
}
