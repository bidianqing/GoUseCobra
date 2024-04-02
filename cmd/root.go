package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionFlag bool
)

func NewMygoCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "mygo",
		Short: "mygo命令行应用程序短介绍",
		Long:  `mygo命令行应用程序详细介绍`,
		Run: func(cmd *cobra.Command, args []string) {
			if versionFlag {
				fmt.Println("Version: 1.0.0")
			} else {
				cmd.Help()
			}
		},
	}

	cmds.ResetFlags()

	cmds.Flags().BoolVarP(&versionFlag, "version", "v", false, "version for mygo")

	cmds.AddCommand(newCmdInit())
	cmds.AddCommand(newCmdRun())

	return cmds
}
