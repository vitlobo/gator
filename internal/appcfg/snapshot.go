package appcfg

import (
	"github.com/vitlobo/gator/internal/gatorsave"
)

func snapshotFromConfig(cfg Config) gatorsave.SaveV1 {
	snap := gatorsave.SaveV1{Version: gatorsave.SaveVersion}
	snap.CurrentUserName = cfg.CurrentUserName
	snap.DBURL = cfg.DBURL

	return snap
}

func ApplySnapshot(cfg *Config, s gatorsave.SaveV1) {
	cfg.CurrentUserName = s.CurrentUserName
	cfg.DBURL = s.DBURL
}