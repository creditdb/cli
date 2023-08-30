package cmd

import "fmt"
func HelpCmd(){
	fmt.Println("Usage: help")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  help - show this help")
	fmt.Println("  select <page> - select a page")
	fmt.Println("  set <key> <value> - set a line")
	fmt.Println("  get <key> - get a line")
	fmt.Println("  getall - get all lines")
	fmt.Println("  del <key> - delete a line")
	fmt.Println("  flush - flush all lines in current page")
	fmt.Println("  ping - ping server")
	fmt.Println("  cls - clear screen")
	fmt.Println("  exit (or Ctrl+C) - exit shell ")
	fmt.Println()
}