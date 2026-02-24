package cmd

import (
	"errors"
	"fmt"
	"grafana-ops/internal/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Config  types.Config
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gops",
	Short: "Grafana Operations CLI",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return loadConfig(cmd)
	},
	Long: `Gops is a tool for advanced operations on Grafana instances, including
	
backup, restore and dashboard management.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grafana-ops/config.yaml)")
	rootCmd.PersistentFlags().String("grafana-token", "", "Grafana API Token")
	rootCmd.PersistentFlags().String("grafana-url", "", "Grafana URL")
	rootCmd.PersistentFlags().String("output", "./_output", "Directory to store backup")
	rootCmd.PersistentFlags().String("loglevel", "info", "Log level")
}

func loadConfig(cmd *cobra.Command) error {
	v := viper.New()

	// ENV (priority 2)
	v.SetEnvPrefix("GO")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	// File (priority 3)
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		v.AddConfigPath(".")
		v.AddConfigPath(home + "/.gops")
		v.SetConfigName("config")
		v.SetConfigType("yaml")
	}

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	// Flags (priority 1)
	if err := v.BindPFlags(cmd.Root().PersistentFlags()); err != nil {
		return err
	}

	if err := v.Unmarshal(&Config); err != nil {
		return fmt.Errorf("error to unmarshal config: %w", err)
	}

	return nil
}
