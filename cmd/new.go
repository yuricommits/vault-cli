package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuricommits/vault-cli/internal/api"
	"github.com/yuricommits/vault-cli/internal/config"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new snippet interactively",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.IsAuthenticated() {
			return fmt.Errorf("not authenticated. Run: vault auth login --token <token>")
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		fmt.Print("Language: ")
		language, _ := reader.ReadString('\n')
		language = strings.TrimSpace(language)

		fmt.Print("Description (optional): ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		fmt.Println("Code (paste code, then type END on a new line):")
		var codeLines []string
		for {
			line, _ := reader.ReadString('\n')
			line = strings.TrimRight(line, "\n")
			if strings.TrimSpace(line) == "END" {
				break
			}
			codeLines = append(codeLines, line)
		}
		code := strings.Join(codeLines, "\n")

		client := api.New()
		var snippet Snippet
		if err := client.Post("/api/snippets", map[string]string{
			"title":       title,
			"language":    language,
			"description": description,
			"code":        code,
		}, &snippet); err != nil {
			return err
		}

		fmt.Printf("✓ Created snippet \"%s\" (%s)\n", snippet.Title, snippet.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
