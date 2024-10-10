package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exenin/profilerz/util"
)

const baseDir = "~/profilerz.d"
const activeProfileDir = "~/.active_profile"

func GetProfilePath(profileName, configType string) string {
	return filepath.Join(util.ExpandPath(baseDir), profileName, configType)
}

func AddProfile(profileName string) error {
	profilePath := filepath.Join(util.ExpandPath(baseDir), profileName)
	if _, err := os.Stat(profilePath); !os.IsNotExist(err) {
		return fmt.Errorf("profile '%s' already exists", profileName)
	}
	return os.MkdirAll(profilePath, os.ModePerm)
}

func ListProfiles() ([]string, error) {
	files, err := os.ReadDir(util.ExpandPath(baseDir))
	if err != nil {
		return nil, err
	}

	var profiles []string
	for _, f := range files {
		if f.IsDir() {
			profiles = append(profiles, f.Name())
		}
	}
	return profiles, nil
}

func SetActiveProfile(profileName string) error {
	profilePath := filepath.Join(util.ExpandPath(baseDir), profileName)
	if _, err := os.Stat(profilePath); os.IsNotExist(err) {
		return fmt.Errorf("profile '%s' does not exist", profileName)
	}

	// Clear current active profile directory
	if err := os.RemoveAll(util.ExpandPath(activeProfileDir)); err != nil {
		return fmt.Errorf("failed to clear active profile: %v", err)
	}

	// Copy profile contents to active profile
	err := util.CopyDir(profilePath, util.ExpandPath(activeProfileDir))
	if err != nil {
		return fmt.Errorf("failed to activate profile: %v", err)
	}

	return nil
}
