package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	printFeed(*feed)
	return nil
}

func printFeed(feed RSSFeed) {
	fmt.Printf(" * Title:       %v\n", feed.Channel.Title)
	fmt.Printf(" * Link:        %v\n", feed.Channel.Link)
	fmt.Printf(" * Description: %v\n", feed.Channel.Description)
	
	fmt.Println(" * Items:")
	for i, item := range feed.Channel.Item {
		fmt.Printf("[%d]:\n", i + 1)
		fmt.Printf("   * Title:       %v\n", item.Title)
		fmt.Printf("   * Link:        %v\n", item.Link)
		fmt.Printf("   * Description: %v\n", item.Description)
		fmt.Printf("   * Published:   %v\n", item.PubDate)
	}
}
