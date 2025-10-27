package gatorsave

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Load(path string) (SaveV1, error) {
	var out SaveV1
	if err := ensureDirFor(path); err != nil {
		return out, fmt.Errorf("ensureDirFor: %w", err)
	}

	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return SaveV1{Version: SaveVersion}, nil // no save yet
		}
		return out, fmt.Errorf("read: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return out, fmt.Errorf("stat: %w", err)
	}
	if info.Size() == 0 {
		return SaveV1{Version: SaveVersion}, nil // treat empty as no data
	}

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&out); err != nil {
		return out, fmt.Errorf("decode: %w", err)
	}

	if out.Version == 0 {
		out.Version = SaveVersion
	}
	
	return out, nil
}