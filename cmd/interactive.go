package cmd

import (
	"fmt"
	"os"

	"github.com/Ronak-jain-afk/palette/internal/tui"
	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Launch interactive palette editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "Palette encountered an error and needs to close.\n")
				os.Exit(1)
			}
		}()

		p := tea.NewProgram(tui.InitialModel(noColor))
		if _, err := p.Run(); err != nil {
			return err
		}
		return nil
	},
}
