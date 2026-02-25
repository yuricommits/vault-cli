package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuricommits/vault-cli/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Vault CLI — manage your code snippets from the terminal",
	Long:  `A CLI for interacting with your Vault snippet manager.`,
}

func Execute() {
	config.Init()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
