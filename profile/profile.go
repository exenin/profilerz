package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"profilerz/config"
	"profilerz/util"
)

const baseDir = "~/.profilerz.d"

func GetProfilePath(profileName, configType string) string {
	return filepath.Join(util.ExpandPath(baseDir), profileName, configType)
}

func AddProfile(profileName string) error {
	profilePath := filepath.Join(util.ExpandPath(baseDir), profileName)
	if _, err := os.Stat(profilePath); !os.IsNotExist(err) {
		return fmt.Errorf("profile '%s' already exists", profileName)
	}

	// Create the default config directories
	// for name, configPath := range config.DefaultConfigs {
	for name := range config.DefaultConfigs {
		path := filepath.Join(profilePath, name)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create default config directory %s: %v", path, err)
		}
	}

	return os.MkdirAll(profilePath, os.ModePerm)
}
func DeleteProfile(profileName string) error {
	profilePath := filepath.Join(util.ExpandPath(baseDir), profileName)
	if _, err := os.Stat(profilePath); os.IsNotExist(err) {
		return fmt.Errorf("profile '%s' does not exist", profileName)
	}

	err := os.RemoveAll(profilePath)
	if err != nil {
		return fmt.Errorf("failed to delete profile %s: %v", profileName, err)
	}

	return nil
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

	// Symlink active profile to their original place
	for name, configPath := range config.DefaultConfigs {
		src := filepath.Join(profilePath, name)
		dst := util.ExpandPath(configPath)
		if err := os.Remove(dst); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove existing file %s: %v", dst, err)
		}
		if err := os.Symlink(src, dst); err != nil {
			return fmt.Errorf("failed to symlink %s to %s: %v", src, dst, err)
		}
	}

	return nil
}
