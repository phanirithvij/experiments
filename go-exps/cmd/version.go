package cmd

import (
	"os"

	"github.com/phanirithvij/experiments/go-exps/experiments/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version info",
	Long:  `Prints the version info along with the build time, OS info, build commit and more`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LogVersionInfo()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
