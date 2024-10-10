package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const baseDir = "~/profilerz.d"
const activeProfileDir = "~/.active_profile"

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}

func initGit(dir string) {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to initialize git: %v", err)
	}
	fmt.Println("Git initialized in profile directory")
}

func addProfile(profileName string) {
	profilePath := filepath.Join(expandPath(baseDir), profileName)
	if _, err := os.Stat(profilePath); !os.IsNotExist(err) {
		fmt.Printf("Profile %s already exists\n", profileName)
		return
	}
	err := os.MkdirAll(profilePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create profile: %v", err)
	}

	// Initialize git for the new profile
	initGit(profilePath)

	fmt.Printf("Profile '%s' created at %s\n", profileName, profilePath)
}

func listProfiles() {
	files, err := ioutil.ReadDir(expandPath(baseDir))
	if err != nil {
		log.Fatalf("Failed to list profiles: %v", err)
	}

	fmt.Println("Available profiles:")
	for _, f := range files {
		if f.IsDir() {
			fmt.Println(" -", f.Name())
		}
	}
}

func setActiveProfile(profileName string) {
	profilePath := filepath.Join(expandPath(baseDir), profileName)
	if _, err := os.Stat(profilePath); os.IsNotExist(err) {
		log.Fatalf("Profile %s does not exist\n", profileName)
	}

	// Clear current active profile directory
	if err := os.RemoveAll(expandPath(activeProfileDir)); err != nil {
		log.Fatalf("Failed to clear active profile: %v", err)
	}

	// Copy profile contents to active profile
	err := copyDir(profilePath, expandPath(activeProfileDir))
	if err != nil {
		log.Fatalf("Failed to activate profile: %v", err)
	}

	fmt.Printf("Profile '%s' is now active\n", profileName)
}

func copyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	if err = os.MkdirAll(dst, os.ModePerm); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				log.Fatal(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func showActiveProfile() {
	files, err := ioutil.ReadDir(expandPath(activeProfileDir))
	if err != nil {
		log.Fatalf("Failed to read active profile directory: %v", err)
	}

	fmt.Println("Currently active profile:")
	for _, f := range files {
		fmt.Println(" -", f.Name())
	}
}

func deleteProfile(profileName string) {
	profilePath := filepath.Join(expandPath(baseDir), profileName)
	if _, err := os.Stat(profilePath); os.IsNotExist(err) {
		log.Fatalf("Profile %s does not exist\n", profileName)
	}

	err := os.RemoveAll(profilePath)
	if err != nil {
		log.Fatalf("Failed to delete profile: %v", err)
	}

	fmt.Printf("Profile '%s' deleted\n", profileName)
}

func main() {
	var rootCmd = &cobra.Command{Use: "profilerz"}

	var addCmd = &cobra.Command{
		Use:   "add [profile name]",
		Short: "Add a new profile",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			addProfile(args[0])
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all profiles",
		Run: func(cmd *cobra.Command, args []string) {
			listProfiles()
		},
	}

	var setCmd = &cobra.Command{
		Use:   "set [profile name]",
		Short: "Set a profile as active",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setActiveProfile(args[0])
		},
	}

	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show active profile",
		Run: func(cmd *cobra.Command, args []string) {
			showActiveProfile()
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete [profile name]",
		Short: "Delete a profile",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteProfile(args[0])
		},
	}

	rootCmd.AddCommand(addCmd, listCmd, setCmd, showCmd, deleteCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
