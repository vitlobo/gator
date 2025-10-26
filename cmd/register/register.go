package register

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
)

func init() {
	core.GetRegisteredCommands().Register("register", handlerRegister)
}

func handlerRegister(state *core.State, command core.Command) error {
	if len(command.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", command.Name)
	}
	name := command.Args[0]

	_, err := state.Db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Printf("User %q already exists\n", name)
		os.Exit(1)
	}

	user, err := state.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := state.Cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func printUser(user database.AppUser) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
	fmt.Printf(" * CreatedAt: %v\n", user.CreatedAt)
	fmt.Printf(" * UpdatedAt: %v\n", user.UpdatedAt)
}