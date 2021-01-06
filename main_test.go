package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelp(t *testing.T) {
	help()
}

func TestParseFlags(t *testing.T) {
	{
		// TODO: Add success case
	}

	{
		// Fail cases
		type pattern struct {
			root     string
			sequence string
			parallel string
		}

		pats := []pattern{
			{},
			{root: "models", sequence: "-1", parallel: "-1"},
		}

		for idx, p := range pats {
			flag.CommandLine.Set("dir", p.root)
			flag.CommandLine.Set("sequence", p.sequence)
			flag.CommandLine.Set("parallel", p.parallel)

			assert.Panics(t, func() {
				parseFlags(true)
			}, fmt.Sprintf("case %d failed", idx))
		}
	}
}
