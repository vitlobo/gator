package reset

import (
	"context"
	"fmt"

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

	fmt.Println("Database reset successfully!")
	return nil
}