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

	// Set ssh key permissions
	if err := SetSshKeyPermissions(profileName); err != nil {
		return fmt.Errorf("failed to set ssh key permissions: %v", err)
	}

	return nil
}

func SetSshKeyPermissions(profileName string) error {
	println("Setting ssh key permissions...")

	sshFolder := GetProfilePath(profileName, "ssh")
	err := os.Chmod(sshFolder, 0700)
	if err != nil {
		return fmt.Errorf("failed to set ssh folder permission to 700: %v", err)
	}

	// Set ssh private key files to 600
	privateKeyFiles, err := filepath.Glob(filepath.Join(sshFolder, "*"))
	if err != nil {
		return fmt.Errorf("failed to list ssh private key files: %v", err)
	}
	for _, f := range privateKeyFiles {
		err = os.Chmod(f, 0600)
		if err != nil {
			return fmt.Errorf("failed to set ssh private key file permission to 600: %v", err)
		}
	}

	// Set ssh public key files to 640
	publicKeyFiles, err := filepath.Glob(filepath.Join(sshFolder, "*.pub"))
	if err != nil {
		return fmt.Errorf("failed to list ssh public key files: %v", err)
	}
	for _, f := range publicKeyFiles {
		err = os.Chmod(f, 0640)
		if err != nil {
			return fmt.Errorf("failed to set ssh public key file permission to 640: %v", err)
		}
	}

	return nil
}
