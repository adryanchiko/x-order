package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/adryanchiko/x-order/service/order-app/pkg/db"
	"github.com/adryanchiko/x-order/service/order-app/pkg/registry"
	"github.com/adryanchiko/x-order/service/order-app/pkg/settings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	// commands
	_ "github.com/adryanchiko/x-order/service/order-app/cmd/order/command"

	// modules
	_ "github.com/adryanchiko/x-order/service/order-app/api/order"
)

func start(config *settings.Settings) error {
	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.Lvl(config.App.Server.LogLevel))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	groupBase := e.Group(config.App.Server.APIBase)

	for _, fn := range registry.ServiceFactories() {
		service := fn(config)
		service.RegisterRoutes(groupBase)
	}

	// Start server
	go func() {
		if err := e.Start(":" + config.App.Server.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}

func main() {
	config, err := settings.Load()
	if err != nil {
		log.Fatal("error loading config", err)
		return
	}

	config.App.Version = fmt.Sprintf("v%d.%d", 0, 1)
	rootCmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: config.App.Name,
		Long:  config.App.Description,
		Run: func(cmd *cobra.Command, args []string) {
			// auto starts the DB pools
			if err := db.Open(config); err != nil {
				log.Fatal(err)
				return
			}

			// start server
			start(config)

			// close connection
			err = db.Close(config)
			if err != nil {
				log.Fatal(err)
				return
			}
		},
	}

	// appends the registered commands
	for _, f := range registry.CommandFactories() {
		// create command
		cmd := f(config)

		// remove registered sub command with same name
		for _, sub := range rootCmd.Commands() {
			if sub.Use == cmd.Use {
				rootCmd.RemoveCommand(sub)
			}
		}

		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		return
	}
}
