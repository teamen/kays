package migrate

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	nestedset "github.com/longbridgeapp/nested-set"
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teamen/kays/internal/apiserver/config"
	"github.com/teamen/kays/internal/apiserver/options"
	"github.com/teamen/kays/internal/apiserver/store/mysql"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

var cfgFile string

const (
	defaultConfigName  = "migrate"
	recommendedHomeDir = ".kays"
)

func NewMigrateCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:          "migrate",
		Short:        "The MIGRATE of the Project `kays`",
		Long:         "The MIGRATE of the Project `kays`",
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

	// fmt.Println("HERE IS MIGRATE RUN")

	db := mysqlStore.DB()

	db.Migrator().DropTable(&v1.Category{})
	db.Migrator().CreateTable(&v1.Category{})
	db.Migrator().CreateTable(&v1.User{})

	node := &v1.Category{
		Title: "产品类别",
		Slug: "	product-type",
	}

	nestedset.Create(db, node, nil)
	// nestedset.Create(db, node, nil)

	// var parent v1.Category
	// db.Model(&v1.Category{}).First(&parent)
	// fmt.Printf("%+v\n\n", parent)

	child1 := &v1.Category{
		Title: "镜框",
		Slug:  "product-type__frame",
		ParentID: sql.NullInt64{
			Int64: node.ID,
			Valid: true,
		},
	}
	fmt.Printf("%+v\n\n", node)

	// fmt.Printf("%+v\n\n", child1)
	// fmt.Printf("%+v\n\n", child2)

	nestedset.Create(db, child1, node)

	child2 := &v1.Category{
		Title: "镜片",
		Slug:  "product-type__lens",
		// the nestedset lib need us to hold the parent id ...
		ParentID: sql.NullInt64{
			Int64: node.ID,
			Valid: true,
		},
	}

	nestedset.Create(db, child2, node)
	// nestedset.MoveTo(db, child2, node, nestedset.MoveDirectionLeft)

	return nil
}
