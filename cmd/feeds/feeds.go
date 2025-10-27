package feeds

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
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

	blue := color.New(color.FgBlue).SprintFunc()

	for _, feed := range feeds {
		fmt.Printf(" * ID:      %s\n", blue(feed.ID))
		fmt.Printf(" * Name:    %s\n", blue(feed.Name))
		fmt.Printf(" * URL:     %s\n", blue(feed.Url))
		fmt.Printf(" * User:    %s\n", blue(feed.Username))
		fmt.Printf(" * Created: %s\n", blue(feed.CreatedAt))
		fmt.Printf(" * Updated: %s\n", blue(feed.UpdatedAt))
		fmt.Println("====================================================")
	}

	return nil
}
