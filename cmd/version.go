package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "打印版本号",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version 1.0.0")
		},
	}

	return cmd
}
