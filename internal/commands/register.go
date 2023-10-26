package commands

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

// Initialize registers the standard commands
func Initialize() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&printCmd{}, "")
}

func DoExecute() {
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
