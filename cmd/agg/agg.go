package agg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
)

func init() {
	core.GetRegisteredCommands().Register("agg", core.MiddlewareLoggedIn(handlerAgg))
}

func handlerAgg(state *core.State, command core.Command, user database.AppUser) error {
	color.Blue("Starting aggregation for %s...", user.Name)

	if len(command.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", command.Name)
	}

	interval, err := time.ParseDuration(command.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	color.Blue("Collecting feeds every %s...", interval)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	scrapeFeeds(state)

	for range ticker.C {
		scrapeFeeds(state)
	}

	return nil
}

func scrapeFeeds(state *core.State) {
	ctx := context.Background()

	feed, err := state.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		color.Red("✗ error selecting next feed: %v", err)
		return
	}

	scrapeFeed(state, feed)
}

func scrapeFeed(state *core.State, feed database.AppFeed) {
	ctx := context.Background()

	feedData, err := state.GatorClient.FetchFeed(ctx, feed.Url)
	if err != nil {
		color.Red("✗ couldn't fetch feed %s: %v", feed.Name, err)
		return
	}

	_, err = state.Db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		color.Red("✗ couldn't mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	if len(feedData.Channel.Item) == 0 {
		color.Yellow("• no posts found for feed %q", feed.Name)
		return
	}

	fmt.Println("====================================================")
	fmt.Println()

	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if pubTime, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  pubTime,
				Valid: true,
			}
		}

		_, err = state.Db.CreatePost(ctx, database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
		})

		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			color.Red("✗ couldn't create post for feed %s: %v", feed.Name, err)
			continue
		}
	}

	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("%s %s %s (%d posts)\n", green("✓"), blue(feed.Name), blue("feed collected"), len(feedData.Channel.Item))
	fmt.Println()
	fmt.Println("====================================================")
}
