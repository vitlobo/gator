package addfeed

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
)

func init() {
	core.GetRegisteredCommands().Register("addfeed", handlerAddFeed)
}

func handlerAddFeed(state *core.State, command core.Command) error {
	if len(command.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", command.Name)
	}
	name := command.Args[0]
	url := command.Args[1]

	ctx := context.Background()

	user, err := state.Db.GetUser(ctx, state.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := state.Db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	color.Blue("Feed added successfully:")
	fmt.Println("====================================================")
	fmt.Println()
	printFeed(feed)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}

func printFeed(feed database.AppFeed) {
	fmt.Printf(" * ID:        %v\n", feed.ID)
	fmt.Printf(" * UserID:    %v\n", feed.UserID)
	fmt.Printf(" * Name:      %v\n", feed.Name)
	fmt.Printf(" * Url:       %v\n", feed.Url)
	fmt.Printf(" * CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt: %v\n", feed.UpdatedAt)
}