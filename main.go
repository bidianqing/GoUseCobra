package main

import (
	"fmt"
	"os"

	"github.com/bidianqing/go-use-cobra/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.OnInitialize(func() {
		fmt.Println("cobra.OnInitialize执行")
	})
	cmd := cmd.NewMygoCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
