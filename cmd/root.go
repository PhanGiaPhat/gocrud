package cmd

import (
	"fmt"
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
	db := database.NewDB()
	conn, err := db.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	wr := repository.NewWager(conn)
	srv := server.NewServer(wr)
	srv.Start()
}
