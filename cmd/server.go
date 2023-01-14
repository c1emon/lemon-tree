/*
Copyright © 2022 clemon
*/
package cmd

import (
	"fmt"
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/controller"
	"github.com/c1emon/lemontree/log"
	"github.com/c1emon/lemontree/persister"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var port int
var dbDriverName string
var dbSourceName string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start lemon tree server",
	Run: func(cmd *cobra.Command, args []string) {
		config.SetConfig(port, dbDriverName, dbSourceName)
		defer func() {
			if err := persister.DisConnect(); err != nil {
				log.GetLogger().Warnf("unable close db: %s", err)
			}
		}()

		e := controller.SingletonEchoFactory()
		e.HTTPErrorHandler = controller.HTTPErrorHandler

		loginG := e.Group("/api/v1/login")
		controller.BuildLogin(loginG)
		e.Start(fmt.Sprintf(":%d", port))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))

	serverCmd.PersistentFlags().StringVar(&dbDriverName, "driver", "postgres", "db driver name")
	viper.BindPFlag("driver", serverCmd.PersistentFlags().Lookup("driver"))

	serverCmd.PersistentFlags().StringVar(&dbSourceName, "source", "host=localhost port=5432 user=postgres dbname=lemon_tree password=123456", "db source")
	viper.BindPFlag("source", serverCmd.PersistentFlags().Lookup("source"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
