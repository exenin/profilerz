package cmd

import (
	"fmt"
	"os"

	_ "github.com/exenin/profilerz/cmd/init"
	_ "github.com/exenin/profilerz/cmd/profile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "profilerz",
	Short: "Profile manager for config directories (AWS, kubectl, DigitalOcean, etc.)",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands (init, profile operations)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(profileCmd)
}
