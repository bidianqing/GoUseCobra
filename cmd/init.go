package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var port string

func newCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init命令",
		Long:  `init命令`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("port ", port)
		},
	}

	cmd.Flags().StringVarP(&port, "port", "p", "", "端口")

	return cmd
}
