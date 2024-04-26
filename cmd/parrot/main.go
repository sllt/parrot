package main

import (
	"fmt"
	"github.com/sllt/parrot/internal/command/create"
	"github.com/sllt/parrot/internal/command/new"
	"github.com/sllt/parrot/internal/command/run"
	"github.com/sllt/parrot/internal/command/upgrade"
	"github.com/sllt/parrot/internal/command/wire"
	"github.com/spf13/cobra"
)

var CmdRoot = &cobra.Command{
	Use:     "parrot",
	Example: "parrot new api demo-api",
}

func init() {
	CmdRoot.AddCommand(new.CmdNew)
	CmdRoot.AddCommand(create.CmdCreate)
	CmdRoot.AddCommand(run.CmdRun)

	CmdRoot.AddCommand(upgrade.CmdUpgrade)
	create.CmdCreate.AddCommand(create.CmdCreateHandler)
	create.CmdCreate.AddCommand(create.CmdCreateService)
	create.CmdCreate.AddCommand(create.CmdCreateRepository)
	create.CmdCreate.AddCommand(create.CmdCreateModel)
	create.CmdCreate.AddCommand(create.CmdCreateAll)

	CmdRoot.AddCommand(wire.CmdWire)
	wire.CmdWire.AddCommand(wire.CmdWireAll)
}

// Execute executes the root command.
func Execute() error {
	return CmdRoot.Execute()
}

func main() {
	err := Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
