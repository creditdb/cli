package cmd

import (
	"context"
	"fmt"

	"github.com/creditdb/go-creditdb"
)

func DelCmd(ctx context.Context, client *creditdb.CreditDB, args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: del <key>")
		fmt.Println()
		return
	}
	err := client.DeleteLine(ctx, args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println()
		return
	}
	fmt.Println("OK")
	fmt.Println()
}
