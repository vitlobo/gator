package gatorsave

import (
	"fmt"
	"os"
	"path/filepath"
)

const SaveVersion    = 1
const configFileName = "gatorconfig.json"

// SaveV1 -
type SaveV1 struct {
	Version     int    `json:"version"`
	CurrentUser string `json:"current_user_name"`
	DB          string `json:"db_url"`
}

func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil { return "", fmt.Errorf("home dir: %w", err) }
	return filepath.Join(home, "Documents", "gator", configFileName), nil
}

func ensureDirFor(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}
	return nil
}