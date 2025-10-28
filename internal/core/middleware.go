package core

import (
	"context"
	"fmt"

	"github.com/vitlobo/gator/internal/database"
)

func MiddlewareLoggedIn(
	handler func(state *State, command Command, user database.AppUser) error,
) func(*State, Command) error {
	return func(state *State, command Command) error {
		if state.Cfg.CurrentUserName == "" {
			return fmt.Errorf("no user logged in - use 'login <name>' first")
		}

		ctx := context.Background()
		user, err := state.Db.GetUser(ctx, state.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find logged-in user %q: %w", state.Cfg.CurrentUserName, err)
		}

		return handler(state, command, user)
	}
}
