package users

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
)

func init() {
	core.GetRegisteredCommands().Register("users", handlerListUsers)
}

func handlerListUsers(state *core.State, command core.Command) error {
	users, err := state.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	fmt.Println("====================================================")
	fmt.Println()

	for _, user := range users {
		if user == state.Cfg.CurrentUserName {
			fmt.Printf("* %s ", user)
			color.Blue("(current)")
			continue
		}
		fmt.Printf("* %s\n", user)
	}

	fmt.Println()
	fmt.Println("====================================================")

	return nil
}