package feeds

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/util"
)

func init() {
	core.GetRegisteredCommands().Register("feeds", handlerListFeeds)
}

func handlerListFeeds(state *core.State, command core.Command) error {
	ctx := context.Background()

	feeds, err := state.Db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		color.Yellow("No feeds found.")
		return nil
	}

	color.New(color.FgBlue).Printf("Found %d feeds:\n", len(feeds))
	fmt.Println("====================================================")
	fmt.Println()
	util.PrintFeeds(feeds)

	return nil
}
