package following

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/util"
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

	feedFollows, err := state.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds for %s: %w", user.Name, err)
	}

	if len(feedFollows) == 0 {
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Printf("%s isn't following any feeds yet.\n", yellow(user.Name))
		return nil
	}

	color.Blue("Feeds followed by %s:", user.Name)
	fmt.Println("====================================================")
	fmt.Println()
	util.PrintFeedsForUser(feedFollows)
	fmt.Println()
	fmt.Println("====================================================")

	return nil
}
