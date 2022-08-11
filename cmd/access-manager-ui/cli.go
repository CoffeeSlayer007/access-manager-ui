package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/roppenlabs/access-manager-ui/internal/config"
	"github.com/roppenlabs/access-manager-ui/internal/server"
	"github.com/roppenlabs/access-manager-ui/pkg/logger"
	"github.com/spf13/cobra"
)

func initCli() *cobra.Command {
	var appCmd = &cobra.Command{
		Use:   "access-manager-ui",
		Short: "Web Interface for access-manager",
	}
	appCmd.AddCommand(startCommand())
	return appCmd
}

func startCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts the service",
		Run: func(cmd *cobra.Command, args []string) {
			configFile := "application"
			if len(args) > 0 {
				configFile = args[0]
			}

			config, err := config.InitConfig(configFile)
			if err != nil {
				panic(fmt.Errorf("error Initializing Config %v", err))
			}

			logger.Init(config.LogLevel)

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			ctx, cancelFn := context.WithCancel(context.Background())

			s := server.New(config)

			go s.Run(ctx)

			<-sigChan

			logger.Info(logger.Format{Message: "Received Signal for Shutdown"})

			cancelFn()

			logger.Info(logger.Format{Message: "Exiting now, Bye!"})
		},
	}

	return startCmd
}
