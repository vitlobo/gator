package reset

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
)

func init() {
	core.GetRegisteredCommands().Register("reset", handlerReset)
}

func handlerReset(state *core.State, command core.Command) error {
	err := state.Db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	color.Blue("Database reset successfully!")
	return nil
}