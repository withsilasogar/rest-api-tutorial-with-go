package cmd

import (
	"context"
	"os"
	"test01/cmd/server"
	"test01/config"
	"test01/provider"
	"time"

	"github.com/sirupsen/logrus"
)

func Execute() {
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background()
	config.LoadEnvironment()

	db, err := config.SetUpDatabase()
	if err != nil {
		logrus.Fatalf("Error setting up database %v", err)
	}

	provider.NewProvider(db, server)
	go func() {
		if err := server.Start(ctx, os.Getenv(config.AppPort)); err != nil {
			logrus.Errorf("Error Starting the server %v", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("Server stopped")

	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Error stopping the server %v", err)
	}
}
