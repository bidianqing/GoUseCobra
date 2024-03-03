package main

import (
	"fmt"
	"os"

	"github.com/bidianqing/go-use-cobra/cmd"
)

func main() {
	cmd := cmd.NewMygoCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
