package cmd

import (
	"context"
	"fmt"

	"github.com/creditdb/go-creditdb"
)

func SetCmd(ctx context.Context, client *creditdb.CreditDB, args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: set <key> <value>")
		fmt.Println()
		return
	}
	err := client.SetLine(ctx, args[1], args[2])
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println()
		return
	}
	fmt.Println("OK")
	fmt.Println()
}
