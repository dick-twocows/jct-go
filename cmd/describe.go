package cmd

import (
	"github.com/spf13/cobra"
)

var (
	describeCmd = &cobra.Command{
		Use:   "describe",
		Short: "Describe ? JCT",
	}
)

func init() {
	rootCmd.AddCommand(describeCmd)
}
