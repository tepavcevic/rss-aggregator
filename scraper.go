package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/tepavcevic/rss-aggregator/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %v duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	//doing first scraping immediately instead of waiting for the ticker
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("Error while getting feeds to fetch: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error while marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error while getting feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		log.Println("Found feed item:", item.Title, " on feed ", feed.Name)
	}
	log.Printf("Feed %s collected, %v items found", feed.Name, len(rssFeed.Channel.Items))
}
