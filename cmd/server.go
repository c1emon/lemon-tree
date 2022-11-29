/*
Copyright © 2022 clemon
*/
package cmd

import (
	"fmt"
	"github.com/c1emon/lemontree/router"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"

	"github.com/spf13/cobra"
)

var port int

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		e := router.EchoFactory()
		//http.ResponseWriter
		e.GET("/", func(c echo.Context) error {
			e.Logger.Print("xxxxx")
			return c.String(http.StatusOK, "Hello, World!")
		})
		e.Start(fmt.Sprintf(":%d", port))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
