package cmd

import (
	"fmt"

	"github.com/Ronak-jain-afk/palette/internal/color"
	"github.com/Ronak-jain-afk/palette/internal/display"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a color palette",
	RunE: func(cmd *cobra.Command, args []string) error {
		mood, _ := cmd.Flags().GetString("mood")
		count, _ := cmd.Flags().GetInt("count")
		scheme, _ := cmd.Flags().GetString("scheme")

		palette := color.GenerateFromMood(mood, count, scheme)

		renderer := display.NewRenderer(noColor, verbose)
		output := display.RenderSwatches(palette, renderer)
		fmt.Println(output)

		return nil
	},
}

func init() {
	generateCmd.Flags().String("mood", "dark", "mood preset (dark, vintage, minimal, nature, pastel)")
	generateCmd.Flags().Int("count", 5, "number of colors (2-10)")
	generateCmd.Flags().String("scheme", "analogous", "harmony scheme (analogous, complementary, triadic, tetradic, monochromatic)")
	generateCmd.Flags().String("base-color", "random", "base color as #RRGGBB or 'random'")
}
