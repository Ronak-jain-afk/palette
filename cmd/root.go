package cmd

import (
	"fmt"
	"os"

	"github.com/Ronak-jain-afk/palette/internal/config"
	"github.com/spf13/cobra"
)

var (
	noColor bool
	cfg     *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "palette",
	Short: "A mood-driven color palette generator for the terminal",
	Long: `Palette generates professional, mood-driven color palettes
directly in your terminal. Use it to quickly explore color schemes,
lock colors you like, and export to CSS, JSON, YAML, or plain hex.`,
}

func Execute() {
	cobra.OnInitialize(func() {
		var err error
		cfg, err = config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %v\n", err)
		}
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable color output")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(configureCmd)
}
