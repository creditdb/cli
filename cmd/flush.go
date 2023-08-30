package cmd

import (
	"context"
	"fmt"

	"github.com/creditdb/go-creditdb"
)

func FlushCmd(ctx context.Context, client *creditdb.CreditDB, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: flush")
		fmt.Println()
		return
	}
	err := client.Flush(ctx)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println()
		return
	}
	fmt.Println("OK")
	fmt.Println()

}
