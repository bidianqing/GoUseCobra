package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var env string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run命令",
	Long:  `run命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run ", env)
		fmt.Println(args)
	},
}

func init() {
	runCmd.Flags().StringVarP(&env, "env", "e", "", "环境变量")
	rootCmd.AddCommand(runCmd)
}
