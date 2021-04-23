package lumecmd

import (
	"fmt"
	"sort"
	"strings"
)

func NewCmdHelp() Command {
	return Command{
		Name:  "help",
		Func:  HelpCmd,
		Use:   "<command>",
		Short: "Show help for a command",
	}
}

func HelpCmd(ctx Context) (int, error) {
	if len(ctx.Args) == 0 {
		fmt.Print(printHelp(commandRegistry))
	} else if len(ctx.Args) >= 1 {
		if cmdHelp, err := printCmdHelp(ctx.Args[0]); err == nil {
			fmt.Print(cmdHelp)
		} else {
			fmt.Print(err)
		}
	}

	return ExitSuccess, nil
}

func printHelp(commands map[string]Command) string {
	var b strings.Builder

	var maxLen, cmdLen int
	var keys []string

	for _, c := range commands {
		keys = append(keys, c.Name)
		cmdLen = len(c.Name)
		if cmdLen > maxLen {
			maxLen = cmdLen
		}
	}

	fmt.Fprintf(&b, "usage:\n  lume <command> [<args...>]")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "\ncommands:")

	sort.Strings(keys)

	for _, k := range keys {
		c := commands[k]
		fmt.Fprintf(&b, "  %-*s    %s\n", maxLen, c.Name, c.Short)
	}

	return b.String()
}

func printCmdHelp(name string) (string, error) {
	var b strings.Builder

	subCmd, ok := commandRegistry[name]

	if !ok {
		return "", fmt.Errorf("unknown commnnd: %s\n", name)
	}

	if subCmd.Use != "" {
		fmt.Fprintf(&b, "usage:\n  lume %s %s\n", subCmd.Name, subCmd.Use)
	} else {
		fmt.Fprintf(&b, "usage:\n  lume %s\n", subCmd.Name)
	}

	if subCmd.Flags != nil {
		out := subCmd.Flags.Output()
		defer subCmd.Flags.SetOutput(out)

		fmt.Fprintln(&b)
		fmt.Fprint(&b, "flags:\n")

		subCmd.Flags.SetOutput(&b)
		subCmd.Flags.PrintDefaults()
	}

	return b.String(), nil
}
