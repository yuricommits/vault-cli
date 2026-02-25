package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuricommits/vault-cli/internal/api"
	"github.com/yuricommits/vault-cli/internal/config"
)

var copyCmd = &cobra.Command{
	Use:   "copy [id]",
	Short: "Copy a snippet's code to clipboard",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !config.IsAuthenticated() {
			return fmt.Errorf("not authenticated. Run: vault auth login --token <token>")
		}

		client := api.New()
		var snippet Snippet
		if err := client.Get("/api/snippets/"+args[0], &snippet); err != nil {
			return err
		}

		if err := copyToClipboard(snippet.Code); err != nil {
			// Fallback: print to stdout
			fmt.Println(snippet.Code)
			return nil
		}

		fmt.Printf("✓ Copied \"%s\" to clipboard\n", snippet.Title)
		return nil
	},
}

func copyToClipboard(text string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		cmd = exec.Command("clip")
	default:
		return fmt.Errorf("unsupported platform")
	}
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
