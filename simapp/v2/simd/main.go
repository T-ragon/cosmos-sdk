package main

import (
	"fmt"
	"os"

	"cosmossdk.io/core/transaction"
	serverv2 "cosmossdk.io/server/v2"
	"cosmossdk.io/simapp/v2"
	"cosmossdk.io/simapp/v2/simd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd[serverv2.AppI[transaction.Tx], transaction.Tx]()
	if err := serverv2.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
