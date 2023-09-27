/*
Copyright © 2022 clemon
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/c1emon/lemontree/internal/server"
	"github.com/c1emon/lemontree/internal/setting"
	"github.com/c1emon/lemontree/pkg/gormx"
	"github.com/c1emon/lemontree/pkg/logx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func listenToSystemSignals(ctx context.Context, s *server.Server) {
	signalChan := make(chan os.Signal, 1)
	sighupChan := make(chan os.Signal, 1)

	signal.Notify(sighupChan, syscall.SIGHUP)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sighupChan:
			// if err := log.Reload(); err != nil {
			// 	fmt.Fprintf(os.Stderr, "Failed to reload loggers: %s\n", err)
			// }
		case sig := <-signalChan:
			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			if err := s.Shutdown(ctx, fmt.Sprintf("System signal: %s", sig)); err != nil {
				fmt.Fprintf(os.Stderr, "Timed out waiting for server to shut down\n")
			}
			return
		}
	}
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start lemon tree server",
	Run: func(cmd *cobra.Command, args []string) {

		defer func() {
			if err := gormx.DisConnect(); err != nil {
				logx.GetLogger().Warnf("unable close db: %s", err)
			}
		}()

		// cfg := setting.New(port, dbDriverName, dbSourceName)
		cfg := setting.GetCfg()
		logx.GetLogger().Infof("%v", cfg)

		gormx.Initialize(cfg.DB.Driver, cfg.DB.Source)
		s, _ := server.Initialize(cfg)
		go listenToSystemSignals(context.Background(), s)
		s.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	cfg := setting.GetCfg()
	serverCmd.PersistentFlags().IntVarP(&cfg.Http.Port, "port", "p", 8080, "server port")
	viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))

	serverCmd.PersistentFlags().StringVar(&cfg.DB.Driver, "driver", "postgres", "db driver name")
	viper.BindPFlag("driver", serverCmd.PersistentFlags().Lookup("driver"))

	serverCmd.PersistentFlags().StringVar(&cfg.DB.Source, "source", "host=10.10.0.70 port=5432 user=postgres dbname=lemon_tree password=123456", "db source")
	viper.BindPFlag("source", serverCmd.PersistentFlags().Lookup("source"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
