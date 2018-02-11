package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passwords",
	Short: "passwords will cracks LBC password hash",
	Long:  `passwords will cracks LBC password hash`,
	Run: func(cmd *cobra.Command, args []string) {
		//do stuff
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
