package engine

import (
	"strings"
)

func Parse(in string) Command {
	fields := strings.Fields(in)

	if len(fields) < 2 {
		return PrintCommand("SYNTAX ERROR: Not Enough Parameters")
	}

	name := fields[0]
	args := fields[1:]

	switch name {
	case "print":
		message := strings.Join(args, " ")

		return PrintCommand(message)
	case "cat":
		arg1 := args[0]
		arg2 := args[1]

		return CatCommand(arg1, arg2)
	default:
		return PrintCommand("ERROR: Unknown Command (" + in + ")")
	}
}
