package main

import (
	"fmt"
	"os"

	"cosmossdk.io/simapp"
	"cosmossdk.io/simapp/simd/cmd"

	svrcmd "github.com/T-ragon/cosmos-sdk/v3/server/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
