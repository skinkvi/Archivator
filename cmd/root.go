package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Simple Archivator",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handErr(err)
	}
}

func handErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
