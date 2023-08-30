package cmd

import (
	"context"
	"fmt"

	"github.com/creditdb/go-creditdb"
)

func GetAllCmd(ctx context.Context, client *creditdb.CreditDB, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: getall")
		fmt.Println()
		return
	}
	lines, err := client.GetAllLines(ctx)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println()
		return
	}
	if len(lines) == 0 {
		fmt.Printf("%+v\n\n", lines)
		return
	}
	for _, line := range lines {
		fmt.Printf("%+v\n\n", line)
	}
}