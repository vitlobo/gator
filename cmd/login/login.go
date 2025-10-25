package login

import (
	"fmt"

	"github.com/vitlobo/gator/internal/core"
)

func init() {
	core.GetRegisteredCommands().Register("login", handlerLogin)
}

func handlerLogin(state *core.State, command core.Command) error {
	if len(command.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", command.Name)
	}
	username := command.Args[0]

	if err := state.Cfg.SetUser(username); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User has been set to %s\n", username)
	return nil
}