package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gobin",
		Short: "gobin let's you upload and download documents from the gobin server",
		Long:  `long description`,
	}

	var cfgFile string
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gobin)")
	cobra.OnInitialize(initConfig(cfgFile))

	return cmd
}

func Execute(command *cobra.Command) {
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig(cfgFile string) func() {
	return func() {
		viper.SetDefault("server", "https://xgob.in")
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)

			viper.SetConfigName(".gobin")
			viper.SetConfigType("env")
			viper.AddConfigPath(home)
		}
		viper.SetEnvPrefix("gobin")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		cobra.CheckErr(err)
	}
}
