package cmd

import (
	"context"
	"fmt"

	"github.com/creditdb/go-creditdb"
)

func PingCmd(ctx context.Context, client *creditdb.CreditDB, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: ping")
		fmt.Println()
		return
	}
	pong, err := client.Ping(ctx)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println()
		return
	}
	fmt.Println(pong)
	fmt.Println()
}
