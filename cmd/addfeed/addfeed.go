package addfeed

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
	"github.com/vitlobo/gator/internal/util"
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
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	color.Blue("Feed created successfully:")
	fmt.Println("====================================================")
	fmt.Println()
	util.PrintAddFeed(feed, user)
	fmt.Println()
	color.Blue("Feed followed successfully:")
	fmt.Println("====================================================")

	return nil
}
