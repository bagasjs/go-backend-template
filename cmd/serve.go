package cmd

import (
	"github.com/bagasjs/go-backend-template/http"
	"github.com/spf13/cobra"
)

func onServe(cmd *cobra.Command, args []string) {
    http.Serve()
}

var serveCmd = &cobra.Command {
    Use: "serve",
    Short: "Serve as an http server",
    Long: "Serve as an http serve default at port 8080",
    Run: onServe,
}
