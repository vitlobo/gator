package browse

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
)

func init() {
	core.GetRegisteredCommands().Register("browse", core.MiddlewareLoggedIn(handlerBrowse))
}

func handlerBrowse(state *core.State, command core.Command, user database.AppUser) error {
	ctx := context.Background()
	limit := 2

	if len(command.Args) > 0 {
		if specifiedLimit, err := strconv.Atoi(command.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := state.Db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to get posts: %w", err)
	}

	if len(posts) == 0 {
		color.Yellow("• No posts found for %s", user.Name)
		return nil
	}

	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("%s %s %d %s %s\n", green("✓"), blue("Found"), len(posts), blue("posts for:"), user.Name)
	fmt.Println("====================================================")
	fmt.Println()

	for _, post := range posts {
		fmt.Printf("%s from %s\n", green(post.PublishedAt.Time.Format("Mon Jan 2")), green(post.FeedName))
		fmt.Printf("--- %s ---\n", blue(post.Title))
		if post.Description.Valid && post.Description.String != "" {
			fmt.Printf("    %v\n", post.Description.String)
		}
		fmt.Printf("%s %s\n", blue("Link:"), post.Url)
		fmt.Println("====================================================")
	}
	fmt.Println()

	return nil
}
