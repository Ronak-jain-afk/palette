package cmd

import (
	"fmt"
	"os"

	"github.com/Ronak-jain-afk/palette/internal/color"
	"github.com/Ronak-jain-afk/palette/internal/export"
	"github.com/spf13/cobra"
)

var (
	exportFormat string
	exportFile   string
	exportClip   bool
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a palette to a file or clipboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		mood := getStringFlag(cmd, "mood", cfg.Defaults.Mood)
		count := getIntFlag(cmd, "count", cfg.Defaults.Count)
		scheme := getStringFlag(cmd, "scheme", cfg.Defaults.Scheme)

		p := color.GenerateFromMood(mood, count, scheme)
		output, err := export.Export(p, exportFormat, true)
		if err != nil {
			return err
		}

		if exportClip {
			if err := export.CopyToClipboard(output); err != nil {
				return fmt.Errorf("clipboard unavailable: %w\npipe output instead: %s", err, output)
			}
			fmt.Println("Copied to clipboard")
			return nil
		}

		if exportFile != "" {
			if err := os.WriteFile(exportFile, []byte(output), 0644); err != nil {
				return err
			}
			fmt.Printf("Wrote %s\n", exportFile)
			return nil
		}

		fmt.Print(output)
		return nil
	},
}

func init() {
	exportCmd.Flags().StringVar(&exportFormat, "format", "json", "output format (json, css, hex, yaml)")
	exportCmd.Flags().StringVar(&exportFile, "output", "", "file path to write")
	exportCmd.Flags().BoolVar(&exportClip, "clipboard", false, "copy to clipboard")
	exportCmd.Flags().String("mood", "dark", "mood preset")
	exportCmd.Flags().Int("count", 5, "number of colors")
	exportCmd.Flags().String("scheme", "analogous", "harmony scheme")
}
