package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuricommits/vault-cli/internal/api"
	"github.com/yuricommits/vault-cli/internal/config"
)

var pushCmd = &cobra.Command{
	Use:   "push [file]",
	Short: "Push a local file as a new snippet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.IsAuthenticated() {
			return fmt.Errorf("not authenticated. Run: vault auth login --token <token>")
		}

		filePath := args[0]
		code, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("could not read file: %w", err)
		}

		title, _ := cmd.Flags().GetString("title")
		language, _ := cmd.Flags().GetString("language")
		description, _ := cmd.Flags().GetString("description")

		if title == "" {
			title = filepath.Base(filePath)
		}
		if language == "" {
			language = detectLanguage(filePath)
		}

		client := api.New()
		var snippet Snippet
		if err := client.Post("/api/snippets", map[string]string{
			"title":       title,
			"language":    language,
			"description": description,
			"code":        string(code),
		}, &snippet); err != nil {
			return err
		}

		fmt.Printf("✓ Pushed \"%s\" as snippet (%s)\n", filePath, snippet.ID)
		return nil
	},
}

func detectLanguage(path string) string {
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	switch ext {
	case "ts", "tsx":
		return "typescript"
	case "js", "jsx":
		return "javascript"
	case "py":
		return "python"
	case "go":
		return "go"
	case "rs":
		return "rust"
	case "java":
		return "java"
	case "sh", "bash":
		return "bash"
	case "sql":
		return "sql"
	case "css":
		return "css"
	case "html":
		return "html"
	default:
		return ext
	}
}

func init() {
	pushCmd.Flags().String("title", "", "Snippet title (default: filename)")
	pushCmd.Flags().String("language", "", "Language (default: auto-detect from extension)")
	pushCmd.Flags().String("description", "", "Snippet description")
	rootCmd.AddCommand(pushCmd)
}
