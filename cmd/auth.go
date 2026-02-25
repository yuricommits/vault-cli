package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yuricommits/vault-cli/internal/config"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your Vault instance",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with a CLI token",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")

		if token == "" {
			return fmt.Errorf("--token is required. Generate one at your Vault settings page.")
		}

		config.SetToken(token)
		if url != "" {
			config.SetBaseURL(url)
		}

		if err := config.Save(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("✓ Authenticated successfully")
		fmt.Printf("  URL: %s\n", config.GetBaseURL())
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set(config.KeyToken, "")
		if err := config.Save(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		fmt.Println("✓ Logged out successfully")
		return nil
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current auth status",
	Run: func(cmd *cobra.Command, args []string) {
		if !config.IsAuthenticated() {
			fmt.Println("✗ Not authenticated")
			fmt.Println("  Run: vault auth login --token <token>")
			return
		}
		fmt.Println("✓ Authenticated")
		fmt.Printf("  URL: %s\n", config.GetBaseURL())
	},
}

func init() {
	authLoginCmd.Flags().String("token", "", "CLI token from Vault settings")
	authLoginCmd.Flags().String("url", "", "Base URL of your Vault instance")
	authCmd.AddCommand(authLoginCmd, authLogoutCmd, authStatusCmd)
	rootCmd.AddCommand(authCmd)
}
