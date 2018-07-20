package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ntindall/logfmt2json/internal"
	"github.com/spf13/cobra"
)

// Main is the entry point to the logfmt2json command line interface.
func Main() {
	rootCmd := &cobra.Command{
		Use:   "logfmt2json",
		Short: "reads logfmt log messages from stdin and prints json to stdout",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			err := internal.Convert(scanner)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
