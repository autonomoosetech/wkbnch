package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wkbnch",
	Short: "A source code generator for SchemaCAN definitions.",
	//Long: ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.wkbnch.yaml")
}

// initConfig reads in config file and env variables if set
func initConfig() {
	if cfgFile != "" {
		// use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// get working directory
		workingDir, err := os.Getwd()
		cobra.CheckErr(err)

		// search config in the working directory with name ".wkbnch" (without extension)
		viper.AddConfigPath(workingDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".wkbnch")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// if a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
