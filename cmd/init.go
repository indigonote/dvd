package cmd

func init() {
	rootCmd.PersistentFlags().StringVarP(&directory, "dir", "d", "", "target directory")
	rootCmd.PersistentFlags().IntVarP(&sequence, "sequence", "s", 0, "a number to specify running machine")
	rootCmd.PersistentFlags().IntVarP(&parallel, "parallel", "p", 1, "max parallelism")
}
