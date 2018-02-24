package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd(version, buildDate string) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of hash-cracker",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hash-cracker version %s - %s\n", version, buildDate)
		},
	}
	return versionCmd
}
