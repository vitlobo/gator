package follow

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
	core.GetRegisteredCommands().Register("follow", core.MiddlewareLoggedIn(handlerFollow))
}

func handlerFollow(state *core.State, command core.Command, user database.AppUser) error {
	if len(command.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", command.Name)
	}
	url := command.Args[0]

	ctx := context.Background()

	feed, err := state.Db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("Feed not found in database: %s", url)
	}

	ffRow, err := state.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	color.Blue("Feed followed successfully:")
	fmt.Println("====================================================")
	fmt.Println()
	util.PrintFeedFollow(ffRow.UserName, ffRow.FeedName)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}
