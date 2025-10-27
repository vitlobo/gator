package follow

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
	core.GetRegisteredCommands().Register("follow", handlerFollow)
}

func handlerFollow(state *core.State, command core.Command) error {
	if len(command.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", command.Name)
	}
	url := command.Args[0]

	ctx := context.Background()

	feed, err := state.Db.GetFeedFromUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("Feed not found in database: %s", url)
	}

	user, err := state.Db.GetUser(ctx, state.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feedFollow, err := state.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed_follow record: %w", err)
	}

	feedFollow.FeedName = feed.Name
	feedFollow.UserName = user.Name

	color.Blue("Feed followed successfully:")
	fmt.Println("====================================================")
	fmt.Println()
	printFeedFollow(feed, user)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}

func printFeedFollow(feed database.GetFeedFromUrlRow, user database.AppUser) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%s %v\n", blue(" * Feed:       "), feed.Name)
	fmt.Printf("%s %v\n", blue(" * URL:        "), feed.Url)
	fmt.Printf("%s %v\n", blue(" * Followed By:"), user.Name)
	fmt.Printf("%s %v\n", blue(" * Feed ID:    "), feed.ID)
}
