package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmdtool",
	Short: "cmd tool",
	Long:  "cmd tool is powerful tool builed by cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmt tool called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
