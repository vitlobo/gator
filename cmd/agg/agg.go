package agg

import (
	"context"
	"fmt"

	"github.com/vitlobo/gator/internal/core"
)

func init() {
	core.GetRegisteredCommands().Register("agg", handlerAgg)
}

func handlerAgg(state *core.State, command core.Command) error {
	err := state.GatorClient.Agg(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed(s): %w", err)
	}
	
	return nil
}