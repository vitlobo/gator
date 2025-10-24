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

	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return SaveV1{Version: SaveVersion}, nil // no save yet
		}
		return out, fmt.Errorf("read: %w", err)
	}
	if len(b) == 0 {
		return SaveV1{Version:  SaveVersion}, nil // treat empty as no data
	}

	if err := json.Unmarshal(b, &out); err != nil {
		return out, fmt.Errorf("unmarshal: %w", err)
	}
	if out.Version == 0 {
		out.Version = SaveVersion
	}
	
	return out, nil
}