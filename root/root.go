package root

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/creditdb/creditsh/cmd"
	"github.com/creditdb/go-creditdb"
	"github.com/spf13/cobra"
)

func CreateRootCommand(client *creditdb.CreditDB, version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "creditsh",
		Short:   "CreditDB Shell",
		Long:    `CLI and Shell for CreditDB`,
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			runShell(client)
		},
	}
	return rootCmd
}

func runShell(client *creditdb.CreditDB) {
	fmt.Printf("Welcome to CreditDB Interactive Shell %v\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Type 'help' for help or 'exit' to exit")
	fmt.Println()
	prompt.New(
		executor(client), completer, prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
			return fmt.Sprintf(" creditdb[%v] >>> ", client.GetCurrentPage()), true
		}), prompt.OptionTitle("CreditDB Shell"),
	).Run()
}

func executor(client *creditdb.CreditDB) func(string) {
	ctx := context.Background()

	return func(in string) {
		args := strings.Fields(in)
		if len(args) == 0 {
			return
		}

		switch args[0] {
		case "help":
			cmd.HelpCmd()
		case "select":
			cmd.SelectCmd(client, args)

		case "set":
			cmd.SetCmd(ctx, client, args)

		case "get":
			cmd.GetCmd(ctx, client, args)

		case "getall":
			cmd.GetAllCmd(ctx, client, args)

		case "del":
			cmd.DelCmd(ctx, client, args)

		case "flush":
			cmd.FlushCmd(ctx, client, args)

		case "ping":
			cmd.PingCmd(ctx, client, args)

		case "cls":
			cmd.ClsCmd()
		case "exit":
			cmd.ExitCmd(ctx, client, "Exiting interactive shell...")
		default:
			fmt.Println("Unknown command:", args[0])
			fmt.Println("Type 'help' for help or 'exit' to exit")
			fmt.Println()
		}
	}

}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}
