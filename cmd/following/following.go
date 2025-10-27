package following

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
)

func init() {
	core.GetRegisteredCommands().Register("following", handlerFollowing)
}

func handlerFollowing(state *core.State, command core.Command) error {
	ctx := context.Background()

	user, err := state.Db.GetUser(ctx, state.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feedsForUser, err := state.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds for %s: %w", user.Name, err)
	}

	if len(feedsForUser) == 0 {
		color.Yellow("%s isn't following any feeds yet.", user.Name)
		return nil
	}

	color.Blue("Feeds followed by %s:", user.Name)
	fmt.Println("====================================================")
	fmt.Println()
	printFeedsForUser(feedsForUser)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}

func printFeedsForUser(feeds []database.GetFeedFollowsForUserRow) {
	blue := color.New(color.FgBlue).SprintFunc()

	for _, feed := range feeds {
		fmt.Printf(" * %s\n", blue(feed.FeedName))
	}
}
