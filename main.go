package main

import (
	"os"

	"github.com/Corwind/supervisorctlbeat/cmd"

	_ "github.com/Corwind/supervisorctlbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
