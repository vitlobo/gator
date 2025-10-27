package gatorsave

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Save(path string, s SaveV1) error {
	if err := ensureDirFor(path); err != nil {
		return fmt.Errorf("ensureDirFor: %w", err)
	}

	s.Version = SaveVersion
	if s.CurrentUserName == "" {
		return errors.New("unable to retrieve current user")
	}

	dir := filepath.Dir(path)
	base := filepath.Base(path)

	tmpFile, err := os.CreateTemp(dir, base+".tmp-*")
	if err != nil {
		return fmt.Errorf("CreateTemp: %w", err)
	}
	tmpPath := tmpFile.Name()

	defer func() {
		tmpFile.Close()
		os.Remove(tmpPath)
	}()

	encoder := json.NewEncoder(tmpFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("fsync: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close: %w", err)
	}

	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("rename: %w", err)
	}

	if err := os.Chmod(path, 0600); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	if d, err := os.Open(dir); err == nil {
		_ = d.Sync()
		_ = d.Close()
	}

	return nil
}