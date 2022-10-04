package apiserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teamen/kays/internal/apiserver/config"
	"github.com/teamen/kays/internal/apiserver/options"
	"github.com/teamen/kays/internal/apiserver/store/mysql"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/token"
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
	viper.SetEnvPrefix("KAYS")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

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

	settings := viper.AllSettings()

	opts := options.NewOptions()
	mapstructure.Decode(settings["mysql"], opts.MySQLOptions)

	config, _ := config.CreateConfigFromOptions(opts)

	// init mysql store
	mysqlStore, _ := mysql.GetMySQLFactoryOr(config.MySQLOptions)
	defer mysqlStore.Close()

	// init JWT
	signingSecret := viper.GetString("jwt_secret")
	token.Init(signingSecret)
	fmt.Println(signingSecret)

	runMode := viper.GetString("run_mode")
	serverAddr := viper.GetString("addr")
	// init gin
	gin.SetMode(runMode)

	// New a new blank Engine instance without any middleware attached.
	g := gin.New()

	loadRouter(g)
	// set route and middleware

	validation.RegisterTranslations()

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

	// e := gin.New()

	return nil
}
