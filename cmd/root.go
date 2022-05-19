package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/PhanGiaPhat/gocrud/pkg/database"
	"github.com/PhanGiaPhat/gocrud/pkg/repository"
	"github.com/PhanGiaPhat/gocrud/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gocrud",
	Short: "gocrud",
	Long:  "gocrud",
	Run:   serve,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocrud.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("gocrud")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.SetEnvPrefix("app")
	for _, cfgName := range viper.AllKeys() {
		viper.BindEnv(cfgName)
	}
}

func serve(cmd *cobra.Command, args []string) {
	db, err := database.NewDB(database.MysqlCfg{
		Username: viper.GetString("db_user"),
		Password: viper.GetString("db_pass"),
		Host:     viper.GetString("db_host"),
		Port:     viper.GetString("db_port"),
		Name:     viper.GetString("db_name"),
		Driver:   viper.GetString("db_driver"),
	})

	if err != nil {
		panic(err)
	}
	defer db.Close()
	m, err := NewMigrate(db.DB())
	if err != nil {
		log.Println("initial database migrate failed...")
	}
	if m != nil {
		if err := m.Up(); err != nil {
			log.Println("database migrate failed...", err)
		} else {
			log.Println("database migrate successful...")
		}
	}

	dbgorm := database.NewDBGORM()
	conn, err := dbgorm.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	wr := repository.NewMessage(conn)
	srv := server.NewServer(wr)
	srv.Start()
}
