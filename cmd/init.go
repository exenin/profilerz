package cmd

import (
	"fmt"
	"os"

	"profilerz/config"
	"profilerz/profile"
	"profilerz/util"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize profilerz by creating default profile with current configs",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize profilerz
		defaultProfile := "default"
		fmt.Println("Initializing profilerz...")

		// Create the default profile
		err := profile.AddProfile(defaultProfile)
		if err != nil {
			fmt.Printf("Failed to create default profile: %v\n", err)
			return
		}

		// Copy existing default configs (AWS, kubectl, etc.) to default profile
		ifSuccess := false
		for name, configPath := range config.DefaultConfigs {
			src := util.ExpandPath(configPath)
			dst := profile.GetProfilePath(defaultProfile, name)
			err := util.CopyDir(src, dst)
			if err != nil {
				// Try to copy it as a file
				err = util.CopyFile(src, dst)
				if err != nil {
					fmt.Printf("Failed to copy %s config: %v\n", name, err)
				} else {
					fmt.Printf("Copied %s config to default profile\n", name)
				}
			} else {
				fmt.Printf("Copied %s config to default profile\n", name)
				ifSuccess = true
			}
		}

		if ifSuccess {
			for name, configPath := range config.DefaultConfigs {
				src := util.ExpandPath(configPath)
				dst := profile.GetProfilePath(defaultProfile, name)
				err := os.RemoveAll(src)
				if err != nil {
					fmt.Printf("Failed to remove original %s config: %v\n", name, err)
				} else {
					err = os.Symlink(dst, src)
					if err != nil {
						fmt.Printf("Failed to set symlink for original %s config: %v\n", name, err)
					} else {
						fmt.Printf("Set symlink for original %s config to default profile\n", name)
					}
				}
			}
		}

		fmt.Println("Profilerz initialization complete.")
	},
}
