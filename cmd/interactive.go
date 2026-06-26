package cmd

import (
	"github.com/Ronak-jain-afk/palette/internal/tui"
	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Launch interactive palette editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(tui.InitialModel(noColor))
		if _, err := p.Run(); err != nil {
			return err
		}
		return nil
	},
}
