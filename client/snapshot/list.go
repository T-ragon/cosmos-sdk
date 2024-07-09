package snapshot

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/T-ragon/cosmos-sdk/client"
	"github.com/T-ragon/cosmos-sdk/server"
)

// ListSnapshotsCmd returns the command to list local snapshots
var ListSnapshotsCmd = &cobra.Command{
	Use:   "list",
	Short: "List local snapshots",
	RunE: func(cmd *cobra.Command, args []string) error {
		viper := client.GetViperFromCmd(cmd)
		snapshotStore, err := server.GetSnapshotStore(viper)
		if err != nil {
			return err
		}
		snapshots, err := snapshotStore.List()
		if err != nil {
			return fmt.Errorf("failed to list snapshots: %w", err)
		}
		for _, snapshot := range snapshots {
			cmd.Println("height:", snapshot.Height, "format:", snapshot.Format, "chunks:", snapshot.Chunks)
		}

		return nil
	},
}
