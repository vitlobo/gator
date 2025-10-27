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

	_, err := state.Db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Printf("User %q already exists\n", name)
		os.Exit(1)
	}

	user, err := state.Db.CreateUser(context.Background(), database.CreateUserParams{
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
	color.New(color.FgBlue).Print(" * ID:        ")
	fmt.Println(user.ID)
	color.New(color.FgBlue).Print(" * Name:      ")
	fmt.Println(user.Name)
	color.New(color.FgBlue).Print(" * CreatedAt: ")
	fmt.Println(user.CreatedAt.String())
	color.New(color.FgBlue).Print(" * UpdatedAt: ")
	fmt.Println(user.UpdatedAt.String())
}
