package cmd

import (
	"errors"

	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/jclebreton/hash-cracker/providers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

// InitRootCmd configure and initialized hash-cracker command
func InitRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hash-cracker [hashes-path] [dictionary-hash]",
		Short: "hash-cracker is a tool to crack cryptographic hash function",
		Long:  `hash-cracker is a tool to crack cryptographic hash function using Providers and Comparators interfaces`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("Requires two args")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			hashes := providers.New(args[0])
			dictionary := providers.New(args[1])
			hasher := &hashers.LbcHash{}
			Run(hashes, dictionary, hasher)
		},
	}

	return rootCmd
}
