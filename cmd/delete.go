package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuricommits/vault-cli/internal/api"
	"github.com/yuricommits/vault-cli/internal/config"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a snippet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.IsAuthenticated() {
			return fmt.Errorf("not authenticated. Run: vault auth login --token <token>")
		}

		force, _ := cmd.Flags().GetBool("force")
		if !force {
			fmt.Printf("Delete snippet %s? [y/N] ", args[0])
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Aborted.")
				return nil
			}
		}

		client := api.New()
		if err := client.Delete("/api/snippets/" + args[0]); err != nil {
			return err
		}

		fmt.Printf("✓ Deleted snippet %s\n", args[0])
		return nil
	},
}

func init() {
	deleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
	rootCmd.AddCommand(deleteCmd)
}
