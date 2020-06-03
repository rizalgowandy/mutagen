package main

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
	"github.com/mutagen-io/mutagen/pkg/mutagenio"
)

func logoutMain(_ *cobra.Command, _ []string) error {
	return mutagenio.Logout()
}

var logoutCommand = &cobra.Command{
	Use:          "logout",
	Short:        "Log out from mutagen.io",
	Args:         cmd.DisallowArguments,
	RunE:         logoutMain,
	SilenceUsage: true,
}

var logoutConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := logoutCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&logoutConfiguration.help, "help", "h", false, "Show help information")
}
