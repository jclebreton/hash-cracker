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
				return errors.New("Requires two args")
			}

			if !isValidHash(args[1]) {
				return errors.New("Invalid LBC hash")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			p := providers.NewDictionaryFromFile(args[0])
			c := &comparators.LBCPassword{}
			c.SetHash(args[1])
			comparators.Compare(c, p)
		},
	}

	return rootCmd
}

func isValidHash(hash string) bool {
	return 56 == len(hash)
}
