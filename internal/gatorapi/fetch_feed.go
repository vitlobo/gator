package gatorapi

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func (c *Client) Agg(ctx context.Context, feedURL string) error {
	feed, err := c.fetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("fetchFeed: %w", err)
	}

	dat, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal indent: %w", err)
	}
	fmt.Println(string(dat))
	return nil
}

func (c *Client) fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")
	req.Header.Set("Accept", "application/xml")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.Copy(io.Discard, resp.Body)
		return &RSSFeed{}, fmt.Errorf("feedURL %s: status %d", feedURL, resp.StatusCode)
	}

	var feed RSSFeed
	decoder := xml.NewDecoder(resp.Body)
	if err := decoder.Decode(&feed); err != nil {
		return &RSSFeed{}, fmt.Errorf("decode: %w", err)
	}

	unescapeFeed(&feed)
	return &feed, nil
}

func unescapeFeed(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
    	feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
}