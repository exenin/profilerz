package cmd

import (
	"fmt"

	"github.com/exenin/profilerz/profile"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles (add, set, list, delete)",
}

var addProfileCmd = &cobra.Command{
	Use:   "add [profile name]",
	Short: "Add a new profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		err := profile.AddProfile(profileName)
		if err != nil {
			fmt.Printf("Error adding profile: %v\n", err)
		} else {
			fmt.Printf("Profile '%s' created.\n", profileName)
		}
	},
}

var listProfilesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profiles, err := profile.ListProfiles()
		if err != nil {
			fmt.Printf("Error listing profiles: %v\n", err)
			return
		}
		fmt.Println("Profiles:")
		for _, p := range profiles {
			fmt.Println(" -", p)
		}
	},
}

var setProfileCmd = &cobra.Command{
	Use:   "set [profile name]",
	Short: "Set a profile as active",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := profile.SetActiveProfile(args[0])
		if err != nil {
			fmt.Printf("Error setting active profile: %v\n", err)
		} else {
			fmt.Printf("Profile '%s' is now active.\n", args[0])
		}
	},
}

func init() {
	profileCmd.AddCommand(addProfileCmd, listProfilesCmd, setProfileCmd)
}
