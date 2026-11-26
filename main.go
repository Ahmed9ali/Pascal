// Package main is the entry point for the Pascal CLI.
package main

import (
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/ahmed/pascal/internal/cmd"
	_ "github.com/ahmed/pascal/internal/dns"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if os.Getenv("PASCAL_PROFILE") != "" {
		go func() {
			slog.Info("Serving pprof at localhost:6060")
			if httpErr := http.ListenAndServe("localhost:6060", nil); httpErr != nil {
				slog.Error("Failed to pprof listen", "error", httpErr)
			}
		}()
	}

	cmd.Execute()
}
