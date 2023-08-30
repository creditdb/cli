package main

import (
	"context"

	"os"
	"os/signal"

	"github.com/creditdb/creditsh/cmd"
	"github.com/creditdb/creditsh/root"
	"github.com/creditdb/go-creditdb"
)

var version = "1.0.0"

func main() {

	client := creditdb.NewClient()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		cmd.ExitCmd(context.Background(), client, "\nReceived Ctrl+C. Exiting interactive shell...")
	}()
	rootCmd := root.CreateRootCommand(client, version)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
