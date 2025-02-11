package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NewFile creates a new config type
func NewFile() IConfig {
	cfg := &File{}

	return cfg
}

// File is a type which describes a config file
type File struct{}

// Load loads the configuration from the configuration file. If the configuration file doesn't exist
// it will return an empty configuration object.
func (c *File) Load() (*Config, error) {
	file, err := c.Location()
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "unable to check if config file exists", err)
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "unable to read config file", err)
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "unable to parse config", err)
	}
	return &cfg, nil
}

// Save saves the given configuration to the configuration file.
func (c *File) Save(cfg *Config) error {
	file, err := c.Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal config", err)
	}
	rhoasCfgDir, err := DefaultDir()
	if err != nil {
		return err
	}
	if _, err = os.Stat(rhoasCfgDir); os.IsNotExist(err) {
		err = os.Mkdir(rhoasCfgDir, 0700)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to save config", err)
	}
	return nil
}

// Remove removes the configuration file.
func (c *File) Remove() error {
	file, err := c.Location()
	if err != nil {
		return err
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return nil
	}
	err = os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

// Location gets the path to the config file
func (c *File) Location() (path string, err error) {
	if rhoasConfig := os.Getenv("RHOASCONFIG"); rhoasConfig != "" {
		path = rhoasConfig
	} else {
		rhoasCfgDir, err := DefaultDir()
		if err != nil {
			return "", nil
		}
		path = filepath.Join(rhoasCfgDir, "config.json")
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

// DefaultDir returns the default parent directory of the config file
func DefaultDir() (string, error) {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", nil
	}
	return filepath.Join(userCfgDir, "rhoas"), nil
}
