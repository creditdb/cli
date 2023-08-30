package cmd

import (
	"context"
	"fmt"

	"github.com/creditdb/go-creditdb"
)

func GetCmd(ctx context.Context, client *creditdb.CreditDB, args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: get <key>")
		fmt.Println()
		return
	}
	value, err := client.GetLine(ctx, args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println()
		return
	}
	fmt.Printf("%+v\n", *value)
	fmt.Println()

}