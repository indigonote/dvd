package cmd

import (
	"fmt"
	"os"

	"github.com/indigonote/dvd/utils"

	"github.com/spf13/cobra"
)

// Root command and flags.
var rootCmd = &cobra.Command{
	Use:   "dvd",
	Short: "Divide a list of directories into smaller chunks.",
	Run: func(cmd *cobra.Command, args []string) {

		dirs, err := utils.Readdir(directory, excludes)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		chunks, err := utils.Chunk(dirs, sequence, parallel)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, c := range chunks {
			fmt.Printf("%s ", c)
		}
	},
}

// Execute from main package.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
