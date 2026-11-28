// Package main is the entry point for the Pascal CLI.
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/Ahmed9ali/Pascal/internal/cmd"
	"github.com/Ahmed9ali/Pascal/internal/envfile"
	_ "github.com/Ahmed9ali/Pascal/internal/dns"
)

func main() {
	if err := envfile.Load(envfile.DefaultPath()); err != nil {
		fmt.Fprintln(os.Stderr, "pascal:", err)
		os.Exit(1)
	}
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
