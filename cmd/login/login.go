package login

import (
	"context"
	"fmt"

	"github.com/fatih/color"
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

	_, err := state.Db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = state.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("coudn't set current user: %w", err)
	}

	color.New(color.FgBlue).Print("Logged in as: ")
	fmt.Println(username)
	return nil
}