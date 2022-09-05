package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/wmarchesi123/dns-check/pkg/dnstools"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test [domain] [requests]",
	Short: "Run a test",
	Long:  `Run a test using default values or set custom values with flags.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		requests, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dnstools.Run(domain, requests)
	},
}
