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

	_, err = state.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed_follow record: %w", err)
	}

	color.Blue("Feed added and followed successfully:")
	fmt.Println("====================================================")
	fmt.Println()
	printFeed(feed)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}

func printFeed(feed database.AppFeed) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%s %v\n", blue(" * ID:       "), feed.ID)
	fmt.Printf("%s %v\n", blue(" * UserID:   "), feed.UserID)
	fmt.Printf("%s %v\n", blue(" * Name:     "), feed.Name)
	fmt.Printf("%s %v\n", blue(" * URL:      "), feed.Url)
	fmt.Printf("%s %v\n", blue(" * CreatedAt:"), feed.CreatedAt)
	fmt.Printf("%s %v\n", blue(" * UpdatedAt:"), feed.UpdatedAt)
}
