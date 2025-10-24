package appcfg

import (
	"github.com/vitlobo/gator/internal/gatorsave"
)

func snapshotFromConfig(cfg *Config) gatorsave.SaveV1 {
	snap := gatorsave.SaveV1{Version: gatorsave.SaveVersion}
	snap.CurrentUser = cfg.CurrentUserName
	snap.DB = cfg.DBURL

	return snap
}

func ApplySnapshot(cfg *Config, s gatorsave.SaveV1) {
	cfg.CurrentUserName = s.CurrentUser
	cfg.DBURL = s.DB
}