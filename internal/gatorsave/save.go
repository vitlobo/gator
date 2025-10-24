package gatorsave

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Write(path string, s SaveV1) error {
	if err := ensureDirFor(path); err != nil {
		return fmt.Errorf("ensureDirFor: %w", err)
	}
	s.Version = SaveVersion
	if s.CurrentUser == "" {
		return errors.New("unable to retrieve current user")
	}

	data, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	return atomicWriteJSON(path, data, 0600)
}

func atomicWriteJSON(path string, data []byte, perm fs.FileMode) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)

	f, err := os.CreateTemp(dir, base+".tmp-*")
	if err != nil { return fmt.Errorf("CreateTemp: %w", err) }
	tmp := f.Name()
	defer func() { _ = f.Close(); _ = os.Remove(tmp) }()

	if _, err := f.Write(data); err != nil { return fmt.Errorf("write: %w", err) }
    if err := f.Sync(); err != nil { return fmt.Errorf("fsync: %w", err) }
    if err := f.Close(); err != nil { return fmt.Errorf("close: %w", err) }

    if err := os.Rename(tmp, path); err != nil { return fmt.Errorf("rename: %w", err) }
    if err := os.Chmod(path, perm); err != nil { return fmt.Errorf("chmod: %w", err) }

	if d, err := os.Open(dir); err == nil {
		_ = d.Sync(); _ = d.Close()
	}
	return nil
}