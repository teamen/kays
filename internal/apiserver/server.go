package apiserver

import (
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

const (
	defaultConfigName  = "apiserver"
	recommendedHomeDir = ".kays"
)

func NewAPIServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "apiserver",
		Short:        "The API of the Project `kays`",
		Long:         "The API of the Project `kays`",
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
		PostRun: func(cmd *cobra.Command, args []string) {
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		RunE: func(cmd *cobra.Command, args []string) error {

			return run()
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s/%s.yaml)", recommendedHomeDir, defaultConfigName))

	return cmd
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("cannot find the home dir:%s", err.Error())
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(defaultConfigName)
	}
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	// viper.SetEnvPrefix("KAYS")

	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf(err.Error())
		} else {
			log.Fatalf(err.Error())
		}
	}

	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
}

func run() error {

	host := viper.GetString("mysql.host")
	fmt.Println(host)
	return nil
}
