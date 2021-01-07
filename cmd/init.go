package cmd

import (
	"os"
)

// Global flags.
var (
	directory string
	excludes  []string
	sequence  int
	parallel  int
)

func init() {
	pwd, err := os.Getwd()

	if err != nil {
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringVarP(&directory, "dir", "d", pwd, "target directory")
	rootCmd.PersistentFlags().IntVarP(&sequence, "sequence", "s", 0, "a number to specify running machine")
	rootCmd.PersistentFlags().IntVarP(&parallel, "parallel", "p", 1, "max parallelism")
	rootCmd.PersistentFlags().StringArrayVarP(&excludes, "exclude", "e", []string{}, "directories to exclude")
}
