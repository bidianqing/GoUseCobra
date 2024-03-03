package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMygoCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "mygo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at https://gohugo.io`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run rootCmd")
		},
	}

	cmds.ResetFlags()

	cmds.AddCommand(newCmdVersion())
	cmds.AddCommand(newCmdRun())

	return cmds
}
