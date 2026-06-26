package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var noColor bool

var rootCmd = &cobra.Command{
	Use:   "palette",
	Short: "A mood-driven color palette generator for the terminal",
	Long: `Palette generates professional, mood-driven color palettes
directly in your terminal. Use it to quickly explore color schemes,
lock colors you like, and export to CSS, JSON, YAML, or plain hex.`,
}

func Execute() {
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
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Launch interactive palette editor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not yet implemented")
	},
}
