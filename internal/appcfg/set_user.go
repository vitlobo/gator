package appcfg

import (
	"fmt"

	"github.com/vitlobo/gator/internal/gatorsave"
)

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return saveUser(*cfg)
}

func saveUser(cfg Config) error {
	path, err := gatorsave.DefaultPath()
	if err != nil { return err }

	snap := snapshotFromConfig(cfg)
	if err := gatorsave.Save(path, snap); err != nil {
		return fmt.Errorf("save failed: %w", err)
	}

	return nil
}