package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuricommits/vault-cli/internal/api"
	"github.com/yuricommits/vault-cli/internal/config"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search your snippets",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.IsAuthenticated() {
			return fmt.Errorf("not authenticated. Run: vault auth login --token <token>")
		}

		client := api.New()
		var snippets []Snippet
		if err := client.Get("/api/search?q="+args[0], &snippets); err != nil {
			return err
		}

		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
			return nil
		}

		fmt.Printf("%-36s  %-12s  %s\n", "ID", "LANGUAGE", "TITLE")
		fmt.Printf("%-36s  %-12s  %s\n", "----", "--------", "-----")
		for _, s := range snippets {
			fmt.Printf("%-36s  %-12s  %s\n", s.ID, s.Language, s.Title)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
