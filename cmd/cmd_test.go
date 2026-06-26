package cmd

import (
	"testing"
)

func TestGenerateFlags(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		mood, _ := generateCmd.Flags().GetString("mood")
		count, _ := generateCmd.Flags().GetInt("count")
		scheme, _ := generateCmd.Flags().GetString("scheme")

		if mood != "dark" {
			t.Errorf("default mood = %q; want %q", mood, "dark")
		}
		if count != 5 {
			t.Errorf("default count = %d; want %d", count, 5)
		}
		if scheme != "analogous" {
			t.Errorf("default scheme = %q; want %q", scheme, "analogous")
		}
	})

	t.Run("output flag", func(t *testing.T) {
		out, _ := generateCmd.Flags().GetString("output")
		if out != "" {
			t.Errorf("default output = %q; want empty", out)
		}
	})
}

func TestExportFlags(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		mood, _ := exportCmd.Flags().GetString("mood")
		count, _ := exportCmd.Flags().GetInt("count")
		scheme, _ := exportCmd.Flags().GetString("scheme")

		if mood != "dark" {
			t.Errorf("default mood = %q; want %q", mood, "dark")
		}
		if count != 5 {
			t.Errorf("default count = %d; want %d", count, 5)
		}
		if scheme != "analogous" {
			t.Errorf("default scheme = %q; want %q", scheme, "analogous")
		}
	})

	t.Run("export format default", func(t *testing.T) {
		if exportFormat != "json" {
			t.Errorf("default exportFormat = %q; want %q", exportFormat, "json")
		}
	})
}

func TestConfigureArgs(t *testing.T) {
	if configureCmd.Args == nil {
		t.Fatal("configureCmd.Args is nil")
	}
	// cobra.ExactArgs(2) should be set
	if configureCmd.Use != "configure <key> <value>" {
		t.Errorf("configure use = %q; want %q", configureCmd.Use, "configure <key> <value>")
	}
}

func TestRootNoColorFlag(t *testing.T) {
	flag := rootCmd.PersistentFlags().Lookup("no-color")
	if flag == nil {
		t.Fatal("--no-color flag not found")
	}
	if flag.DefValue != "false" {
		t.Errorf("--no-color default = %q; want %q", flag.DefValue, "false")
	}
}
