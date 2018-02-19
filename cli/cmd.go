package cli

import (
	"runtime"

	"github.com/jclebreton/hash-cracker/infrastructures/comparators"
	"github.com/jclebreton/hash-cracker/infrastructures/generators"
	"github.com/jclebreton/hash-cracker/infrastructures/progress"
	"github.com/jclebreton/hash-cracker/infrastructures/readers"
	"github.com/jclebreton/hash-cracker/usecases"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var randomize bool

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetLevel(logrus.InfoLevel)
}

// InitRootCmd configure and initialized hash-cracker command
func InitRootCmd(hasher comparators.Comparator) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hash-cracker [hashes-path] [dictionary-hash]",
		Short: "hash-cracker is a tool to crack cryptographic hash function",
		Long:  `hash-cracker is a tool to crack cryptographic hash function using Providers and Comparators interfaces`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			runCmd(args[0], args[1], hasher)
		},
	}
	rootCmd.Flags().BoolVarP(&randomize, "generate", "g", false, "generate passwords from dictionary")
	return rootCmd
}

func runCmd(hashPath, dictionaryPath string, hasher comparators.Comparator) {
	logrus.Infof("%d logical CPUs", runtime.NumCPU())
	if randomize {
		logrus.Info("passwords dictionary generation enable")
	}

	// Build dictionary provider
	dictionaryProvider := readers.DictionaryReader{
		ProgressBar:        progress.NewProgressBar("Dictionary"),
		DictionaryProvider: readers.NewTextFileReader(dictionaryPath),
		PasswordsGenerator: &generators.Basic{},
	}

	// Build hashes provider
	hashesProvider := readers.HashesReader{
		ProgressBarHashes:  progress.NewProgressBar("Hashes"),
		ProgressBarCracked: progress.NewProgressBar("Cracked"),
		HashesProvider:     readers.NewTextFileReader(hashPath),
	}

	// Run
	crackHashesHandler := &usecases.CrackHashesUsingDictionaryHandler{
		HashComparator:    hasher,
		DictionaryReader:  dictionaryProvider,
		HashesReader:      hashesProvider,
		ProgressBarPooler: &progress.CheggaaaPool{},
	}
	crackHashesHandler.Handle(runtime.NumCPU(), randomize)
}
