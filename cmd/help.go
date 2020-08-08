package lumecmd

import (
	"flag"
	"fmt"
)

func init() {
	var cmdName string = "help"
	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

	RegisterCommand(cmdName, Command{
		Func:  HelpCmd,
		Flags: fs,
		Use:   "<command>",
		Short: "Show help for a command",
	})
}

func HelpCmd(args CmdArgs) (int, error) {
	argv := args.Flags.Args()

	if len(argv) == 0 {
		printHelp(commandRegistry)
	} else if len(argv) >= 1 {
		subCmd, ok := commandRegistry[argv[0]]
		if !ok {
			fmt.Printf("unknown command: %s\n", argv[0])
			return 1, nil
		}

		if subCmd.Use != "" {
			fmt.Printf("usage:\n  lume %s %s\n", subCmd.Name, subCmd.Use)
			fmt.Println()
		}

		fmt.Print("flags:\n")
		subCmd.Flags.PrintDefaults()
	}

	return 0, nil
}

func printHelp(commands map[string]Command) {
	var maxLen, cmdLen int
	for _, c := range commands {
		cmdLen = len(c.Name)
		if cmdLen > maxLen {
			maxLen = cmdLen
		}
	}

	fmt.Printf("usage:\n  lume <command> [<args...>]")
	fmt.Println()

	fmt.Println("\ncommands:")
	for _, c := range commands {
		fmt.Printf("  %-*s    %s\n", maxLen, c.Name, c.Short)
	}
}
