package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Readdir returns a list of flat directories from root path.
func readdir(root string) ([]string, error) {
	files := []string{}

	abs, err := filepath.Abs(root)

	if err != nil {
		return files, err
	}

	info, err := ioutil.ReadDir(abs)

	if err != nil {
		return files, err
	}

	// The first record represents the module root.
	files = append(files, gopath(root, false))

	for _, d := range info {
		if !d.IsDir() {
			continue
		}

		files = append(files, gopath(fmt.Sprintf("%s/%s", root, d.Name()), true))
	}

	return files, nil
}

// Convert to Go flavored path.
//
// e.g.
// - no wildcard: /somedir --> /somedir
// - wildcard:    /somedir --> /somedir/...
func gopath(base string, wildcard bool) string {
	if wildcard {
		return fmt.Sprintf("./%s/...", base)
	}

	return fmt.Sprintf("./%s", base)
}

// Chunk returns a filtered list based on sequence number and max parallelism.
//
// - sequence: set machine number, starting from 0
// - maxparallel: set allowed parallelism, starting from 1
func chunk(given []string, sequence, maxparallel int) ([]string, error) {
	files := []string{}

	if len(given) == 0 {
		return files, errors.New("given list is empty")
	}

	if sequence < 0 || maxparallel < 1 {
		return files, errors.New("sequence must be greater than 0, and max parallism must be greater than 1")
	}

	if sequence+1 > maxparallel {
		return files, fmt.Errorf("sequence (%d+1) must not exceed max parallelism (%d)", sequence, maxparallel)
	}

	total := len(given)

	if maxparallel > total {
		return files, fmt.Errorf("max parallelism (%d) must not exceed target size (%d)", maxparallel, total)
	}

	var job int

	// Calculate how many records (jobs) each sequence should handle.
	if total%maxparallel == 0 {
		job = int(total / maxparallel)
	} else {
		job = int(total/maxparallel) + 1
	}

	from := sequence * job
	to := from + job

	if total < to {
		return given[from:], nil
	}

	return given[from:to], nil
}
