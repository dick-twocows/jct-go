package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	jct_version string = "v0.1"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of JCT",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(jct_version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
