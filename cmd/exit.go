package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/creditdb/go-creditdb"
)

func ExitCmd(ctx context.Context, client *creditdb.CreditDB, message string) {
	fmt.Println(message)
	fmt.Println()
	client.Close(ctx)
	os.Exit(0)
}

