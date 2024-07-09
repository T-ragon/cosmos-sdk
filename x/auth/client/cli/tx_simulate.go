package cli

import (
	"strings"

	"github.com/spf13/cobra"

	authclient "cosmossdk.io/x/auth/client"

	"github.com/T-ragon/cosmos-sdk/client"
	"github.com/T-ragon/cosmos-sdk/client/flags"
	"github.com/T-ragon/cosmos-sdk/client/tx"
)

// GetSimulateCmd returns a command that simulates whether a transaction will be
// successful.
func GetSimulateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simulate /path/to/unsigned-tx.json --from keyname",
		Short: "Simulate the gas usage of a transaction",
		Long: strings.TrimSpace(`Simulate whether a transaction will be successful:

- if successful, the simulation result is printed, which includes the gas
  consumption, message response data, and events emitted;
- if unsuccessful, the error message is printed.

The user must provide the path to a JSON-encoded unsigned transaction, typically
generated by any transaction command with the --generate-only flag. It should
look like below. Note that the "signer_infos" and "signatures" fields are left
empty; they will be auto-populated by dummy data for simulation purpose.

{
  "body": {
    "messages": [
      {
        "@type": "/cosmos.bank.v1beta1.MsgSend",
        "from_address": "cosmos1...",
        "to_address": "cosmos1...",
        "amount": [
          {
            "denom": "utoken",
            "amount": "12345"
          }
        ]
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    },
    "tip": null
  },
  "signatures": []
}

The --from flag is mandatory, as the signer account's correct sequence number is
necessary for simulation.
`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf, err := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			txf, err = txf.Prepare(clientCtx)
			if err != nil {
				return err
			}

			stdTx, err := authclient.ReadTxFromFile(clientCtx, args[0])
			if err != nil {
				return err
			}

			simRes, _, err := tx.CalculateGas(clientCtx, txf, stdTx.GetMsgs()...)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(simRes)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
