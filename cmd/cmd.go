package cmd

import (
	"errors"
	"runtime"

	"github.com/jclebreton/hash-cracker/dictionaries"
	"github.com/jclebreton/hash-cracker/hashers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ramdomize bool

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetLevel(logrus.InfoLevel)
}

// InitRootCmd configure and initialized hash-cracker command
func InitRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hash-cracker [hashes-path] [dictionary-hash]",
		Short: "hash-cracker is a tool to crack cryptographic hash function",
		Long:  `hash-cracker is a tool to crack cryptographic hash function using Providers and Comparators interfaces`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgs(args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			runCmd(args)
		},
	}
	rootCmd.Flags().BoolVarP(&ramdomize, "random", "r", false, "randomize passwords dictionary")
	return rootCmd
}

func checkArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("Requires two args")
	}
	return nil
}

func runCmd(args []string) {
	logrus.Infof("%d logical cpu", runtime.NumCPU())
	if ramdomize {
		logrus.Info("randomize passwords dictionary enable")
	}
	hashes := dictionaries.New(args[0])
	dictionary := dictionaries.New(args[1])
	hasher := &hashers.Sha1WithSalt{}
	Run(hashes, dictionary, hasher, runtime.NumCPU(), ramdomize)
}
