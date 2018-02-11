package cmd

import (
	"errors"

	"github.com/jclebreton/hash-cracker/comparators"
	"github.com/jclebreton/hash-cracker/providers"
	"github.com/spf13/cobra"
)

// InitRootCmd configure and initialized hash-cracker command
func InitRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hash-cracker [dictionary-path] [lbc-hash]",
		Short: "hash-cracker is a tool to crack cryptographic hash function",
		Long:  `hash-cracker is a tool to crack cryptographic hash function using Providers and Comparators interfaces`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("requires two args")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			c := &comparators.LBCPassword{Hash: args[1]}
			p := providers.NewDictionaryFromFile(args[0])
			CrackHash(c, p)
		},
	}

	return rootCmd
}
