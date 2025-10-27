package register

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
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

	ctx := context.Background()

	_, err := state.Db.GetUser(ctx, name)
	if err == nil {
		fmt.Printf("User %q already exists\n", name)
		os.Exit(1)
	}

	user, err := state.Db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := state.Cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	color.Blue("User created successfully:")
	fmt.Println("====================================================")
	fmt.Println()
	printUser(user)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}

func printUser(user database.AppUser) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf(" * ID:        %s\n", blue(user.ID))
	fmt.Printf(" * Name:      %s\n", blue(user.Name))
	fmt.Printf(" * CreatedAt: %s\n", blue(user.CreatedAt))
	fmt.Printf(" * UpdatedAt: %s\n", blue(user.UpdatedAt))
}
