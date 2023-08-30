package cmd

import (
	"fmt"
	"strconv"

	"github.com/creditdb/go-creditdb"
)

func SelectCmd(client *creditdb.CreditDB, args []string){
	if len(args) != 2 {
		fmt.Println("Usage: select <page>")
		fmt.Println()
		return
	}
	num, err := strconv.ParseUint(args[1], 10, 0)
	if err != nil {
		fmt.Println("Usage: select <page<unsigned integer>>")
		fmt.Println()
		return
	}
	client.WithPage(uint(num))
	fmt.Println("Switched to page: ", num)
	fmt.Println()
}