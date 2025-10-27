package util

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/vitlobo/gator/internal/database"
)

func PrintAddFeed(feed database.AppFeed, user database.AppUser) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%s %v\n", blue(" * ID:       "), feed.ID)
	fmt.Printf("%s %v\n", blue(" * Name:     "), feed.Name)
	fmt.Printf("%s %v\n", blue(" * URL:      "), feed.Url)
	fmt.Printf("%s %v\n", blue(" * User:     "), user.Name)
	fmt.Printf("%s %v\n", blue(" * CreatedAt:"), feed.CreatedAt)
	fmt.Printf("%s %v\n", blue(" * UpdatedAt:"), feed.UpdatedAt)
}

func PrintFeeds(feeds []database.GetFeedsRow) {
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
}

func PrintFeedFollow(username, feedname string) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("%s %v\n", blue(" * Feed:       "), feedname)
	fmt.Printf("%s %v\n", blue(" * Followed By:"), username)
}

func PrintFeedsForUser(feedFollows []database.GetFeedFollowsForUserRow) {
	blue := color.New(color.FgBlue).SprintFunc()

	for _, ff := range feedFollows {
		fmt.Printf(" * %s\n", blue(ff.FeedName))
	}
}

func PrintUser(user database.AppUser) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf(" * ID:        %s\n", blue(user.ID))
	fmt.Printf(" * Name:      %s\n", blue(user.Name))
	fmt.Printf(" * CreatedAt: %s\n", blue(user.CreatedAt))
	fmt.Printf(" * UpdatedAt: %s\n", blue(user.UpdatedAt))
}
