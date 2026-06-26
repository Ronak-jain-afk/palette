package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	noColor  bool
	verbose  bool
)

var rootCmd = &cobra.Command{
	Use:   "palette",
	Short: "A mood-driven color palette generator for the terminal",
	Long: `Palette generates professional, mood-driven color palettes
directly in your terminal. Use it to quickly explore color schemes,
lock colors you like, and export to CSS, JSON, YAML, or plain hex.`,
}

func Execute() {
	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable color output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(configureCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := os.UserConfigDir()
		if err == nil {
			viper.AddConfigPath(filepath.Join(configDir, "palette"))
		}
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}
	viper.SetEnvPrefix("PALETTE")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Launch interactive palette editor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not yet implemented")
	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the current palette",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not yet implemented")
	},
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Set default preferences",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not yet implemented")
	},
}
