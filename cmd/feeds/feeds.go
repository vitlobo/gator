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
	feeds, err := state.Db.GetFeeds(context.Background())
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

	for _, feed := range feeds {
		color.New(color.FgBlue).Print(" * ID:      ")
		fmt.Println(feed.ID.String())
		color.New(color.FgBlue).Print(" * Name:    ")
		fmt.Println(feed.Name)
		color.New(color.FgBlue).Print(" * URL:     ")
		fmt.Println(feed.Url)
		color.New(color.FgBlue).Print(" * User:    ")
		fmt.Println(feed.Username)
		color.New(color.FgBlue).Print(" * Created: ")
		fmt.Println(feed.CreatedAt.String())
		color.New(color.FgBlue).Print(" * Updated: ")
		fmt.Println(feed.UpdatedAt.String())
		fmt.Println()
		fmt.Println("====================================================")
	}

	return nil
}
