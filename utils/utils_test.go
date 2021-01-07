package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReaddir(t *testing.T) {
	{
		// Succeeds on the existing directory, in relative path.
		pwd, err := os.Getwd()

		require.Nil(t, err)

		abs, err := filepath.Abs(fmt.Sprintf("%s/test", pwd))

		require.Nil(t, err)

		got, err := Readdir(abs)

		assert.Nil(t, err)
		assert.Greater(t, len(got), 1)
	}

	{
		// Fails on non-existent directory.
		got, err := Readdir("whatever-non-existent")

		assert.NotNil(t, err)
		assert.Equal(t, []string{}, got)
	}
}

func TestGopath(t *testing.T) {
	assert.Equal(t, "./hello", gopath("hello", false))
	assert.Equal(t, "./hello/...", gopath("hello", true))
}

func TestChunk(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			given       []string
			sequence    int
			maxparallel int
			expected    []string
		}

		pats := []pattern{
			{
				given:       []string{"a"},
				sequence:    0,
				maxparallel: 1,
				expected:    []string{"a"},
			},
			{
				given:       []string{"a", "b"},
				sequence:    0,
				maxparallel: 1,
				expected:    []string{"a", "b"},
			},
			{
				// Ensure the first record is returned on 2 parallel runs.
				given:       []string{"a", "b"},
				sequence:    0,
				maxparallel: 2,
				expected:    []string{"a"},
			},
			{
				// Ensure the last record is returned on 2 parallel runs.
				given:       []string{"a", "b"},
				sequence:    1,
				maxparallel: 2,
				expected:    []string{"b"},
			},
			{
				// Ensure the fist record is returned on 2 parallel runs.
				given:       []string{"a", "b", "c"},
				sequence:    0,
				maxparallel: 2,
				expected:    []string{"a", "b"},
			},
			{
				// Ensure the last 2 records are returned on 2 parallel runs.
				given:       []string{"a", "b", "c"},
				sequence:    1,
				maxparallel: 2,
				expected:    []string{"c"},
			},
			{
				// Complex) 1st sequence.
				given:       []string{"a", "b", "c", "d", "e"},
				sequence:    0,
				maxparallel: 3,
				expected:    []string{"a", "b"},
			},
			{
				// Complex) 2nd sequence.
				given:       []string{"a", "b", "c", "d", "e"},
				sequence:    1,
				maxparallel: 3,
				expected:    []string{"c", "d"},
			},
			{
				// Complex) the last sequence.
				given:       []string{"a", "b", "c", "d", "e"},
				sequence:    2,
				maxparallel: 3,
				expected:    []string{"e"},
			},
		}

		for idx, p := range pats {
			got, err := Chunk(p.given, p.sequence, p.maxparallel)

			assert.Nilf(t, err, fmt.Sprintf("case %d", idx))
			assert.Equalf(t, p.expected, got, fmt.Sprintf("case %d", idx))
		}
	}

	{
		// Fail cases
		type pattern struct {
			given       []string
			sequence    int
			maxparallel int
		}

		pats := []pattern{
			{},
			{
				given:       []string{"a"},
				sequence:    -1, // error
				maxparallel: 1,
			},
			{
				given:       []string{"a"},
				sequence:    1,
				maxparallel: 0, // error
			},
			{
				given:       []string{"a"},
				sequence:    1, // error
				maxparallel: 1,
			},
			{
				given:       []string{"a"},
				sequence:    1,
				maxparallel: 2, // error
			},
			{
				given:       []string{"a"},
				sequence:    1,
				maxparallel: 2, // error
			},
		}

		for idx, p := range pats {
			got, err := Chunk(p.given, p.sequence, p.maxparallel)

			assert.NotNilf(t, err, spew.Sdump(p))
			assert.Equalf(t, []string{}, got, fmt.Sprintf("case %d", idx))
		}
	}
}
