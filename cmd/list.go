package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuricommits/vault-cli/internal/api"
	"github.com/yuricommits/vault-cli/internal/config"
)

type Snippet struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Code        string `json:"code"`
	CreatedAt   string `json:"createdAt"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your snippets",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.IsAuthenticated() {
			return fmt.Errorf("not authenticated. Run: vault auth login --token <token>")
		}

		client := api.New()
		var snippets []Snippet
		if err := client.Get("/api/snippets", &snippets); err != nil {
			return err
		}

		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
			return nil
		}

		fmt.Printf("%-36s  %-12s  %s\n", "ID", "LANGUAGE", "TITLE")
		fmt.Printf("%-36s  %-12s  %s\n", "----", "--------", "-----")
		for _, s := range snippets {
			desc := s.Description
			if len(desc) > 50 {
				desc = desc[:50] + "..."
			}
			fmt.Printf("%-36s  %-12s  %s\n", s.ID, s.Language, s.Title)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
