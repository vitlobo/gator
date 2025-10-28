package follow

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
	"github.com/vitlobo/gator/internal/util"
)

func init() {
	core.GetRegisteredCommands().Register("unfollow", core.MiddlewareLoggedIn(handlerUnfollow))
}

func handlerUnfollow(state *core.State, command core.Command, user database.AppUser) error {
	if len(command.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", command.Name)
	}
	url := command.Args[0]

	ctx := context.Background()
	feed, err := state.Db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = state.Db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	color.Blue("Successfully unfollowed:")
	fmt.Println("====================================================")
	fmt.Println()
	util.PrintDeleteFeedFollow(feed)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}
