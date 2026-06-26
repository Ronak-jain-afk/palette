package cmd

import (
	"fmt"

	"github.com/Ronak-jain-afk/palette/internal/config"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure <key> <value>",
	Short: "Set a config default",
	Long: `Set persistent default values for palette commands.

Keys:
  defaults.mood     mood preset (dark, vintage, minimal, nature, pastel)
  defaults.count    number of colors (2-10)
  defaults.scheme   harmony scheme (analogous, complementary, triadic, tetradic, monochromatic)`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, value := args[0], args[1]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		switch key {
		case "defaults.mood":
			valid := map[string]bool{"dark": true, "vintage": true, "minimal": true, "nature": true, "pastel": true}
			if !valid[value] {
				return fmt.Errorf("invalid mood: %s (dark, vintage, minimal, nature, pastel)", value)
			}
			cfg.Defaults.Mood = value
		case "defaults.count":
			var n int
			if _, err := fmt.Sscanf(value, "%d", &n); err != nil || n < 2 || n > 10 {
				return fmt.Errorf("invalid count: %s (2-10)", value)
			}
			cfg.Defaults.Count = n
		case "defaults.scheme":
			valid := map[string]bool{"analogous": true, "complementary": true, "triadic": true, "tetradic": true, "monochromatic": true}
			if !valid[value] {
				return fmt.Errorf("invalid scheme: %s (analogous, complementary, triadic, tetradic, monochromatic)", value)
			}
			cfg.Defaults.Scheme = value
		default:
			return fmt.Errorf("unknown key: %s", key)
		}

		if err := config.Save(cfg); err != nil {
			return err
		}

		fmt.Printf("Set %s = %s\n", key, value)
		return nil
	},
}

func init() {
	configureCmd.Flags().Bool("global", false, "apply to all users (unused)")
	configureCmd.Flags().MarkHidden("global")
}
