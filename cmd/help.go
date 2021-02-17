package lumecmd

import (
	"flag"
	"fmt"
	"sort"
)

func NewCmdHelp() Command {
	return Command{
		Name: "help",
		Func: HelpCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("help", flag.ExitOnError)

			return fs
		}(),
		Use:   "<command>",
		Short: "Show help for a command",
	}
}

func HelpCmd(args CmdArgs) (int, error) {
	argv := args.Flags.Args()

	if len(argv) == 0 {
		printHelp(commandRegistry)
	} else if len(argv) >= 1 {
		printCmdHelp(argv[0])
	}

	return ExitSuccess, nil
}

func printHelp(commands map[string]Command) {
	var maxLen, cmdLen int
	var keys []string

	for _, c := range commands {
		keys = append(keys, c.Name)
		cmdLen = len(c.Name)
		if cmdLen > maxLen {
			maxLen = cmdLen
		}
	}

	fmt.Printf("usage:\n  lume <command> [<args...>]")
	fmt.Println()
	fmt.Println("\ncommands:")

	sort.Strings(keys)
	for _, k := range keys {
		c := commands[k]
		fmt.Printf("  %-*s    %s\n", maxLen, c.Name, c.Short)
	}
}

func printCmdHelp(name string) error {
	subCmd, ok := commandRegistry[name]
	if !ok {
		return fmt.Errorf("unknown commnnd: %s\n", name)
	}

	if subCmd.Use != "" {
		fmt.Printf("usage:\n  lume %s %s\n", subCmd.Name, subCmd.Use)
		fmt.Println()
	}

	fmt.Print("flags:\n")
	subCmd.Flags.PrintDefaults()

	return nil
}
