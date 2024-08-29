package engine

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
	"scylla/pkg/config"
	"syscall"
	"time"
)

func StartServerWithGracefulShutdown(app *fiber.App) {
	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Run the server in a goroutine so that it doesn't block
	go func() {
		if err := app.Listen(":" + config.Get().Server.Port); err != nil {
			panic(err)
		}
	}()

	// Wait for an interrupt signal
	<-quit
	log.Info("Shutting down server...")

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		panic(err)
	}

	log.Info("Server exited gracefully")
}
