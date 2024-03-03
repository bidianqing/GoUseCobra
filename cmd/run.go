package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var env string

func newCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run命令",
		Long:  `run命令`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run ", env)
			fmt.Println(args)
		},
	}

	cmd.Flags().StringVarP(&env, "env", "e", "", "环境变量")

	return cmd
}
