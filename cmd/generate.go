package cmd

import (
	"fmt"

	"github.com/Ronak-jain-afk/palette/internal/color"
	"github.com/Ronak-jain-afk/palette/internal/display"
	"github.com/Ronak-jain-afk/palette/internal/export"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a color palette",
	RunE: func(cmd *cobra.Command, args []string) error {
		mood, _ := cmd.Flags().GetString("mood")
		count, _ := cmd.Flags().GetInt("count")
		scheme, _ := cmd.Flags().GetString("scheme")
		outFmt, _ := cmd.Flags().GetString("output")

		palette := color.GenerateFromMood(mood, count, scheme)

		if outFmt != "" {
			out, err := export.Export(palette, outFmt, true)
			if err != nil {
				return err
			}
			fmt.Print(out)
			return nil
		}

		renderer := display.NewRenderer(noColor)
		out := display.RenderSwatches(palette, renderer)
		fmt.Println(out)

		return nil
	},
}

func init() {
	generateCmd.Flags().String("mood", "dark", "mood preset (dark, vintage, minimal, nature, pastel)")
	generateCmd.Flags().Int("count", 5, "number of colors (2-10)")
	generateCmd.Flags().String("scheme", "analogous", "harmony scheme (analogous, complementary, triadic, tetradic, monochromatic)")
	generateCmd.Flags().String("output", "", "export format instead of display (json, css, hex, yaml)")
}
