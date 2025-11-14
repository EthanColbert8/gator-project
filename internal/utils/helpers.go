package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/EthanColbert8/gator-project/internal/database"
	"github.com/google/uuid"
)

func validateNumArgs(cmd Command, numArgs int) error {
	if len(cmd.Args) != numArgs {
		return fmt.Errorf("got %d args, expected %d for command %s", len(cmd.Args), numArgs, cmd.Name)
	}
	return nil
}

func FollowFeed(db *database.Queries, username, url string) (database.CreateFeedFollowRow, error) {
	user, err := db.GetUser(context.Background(), username)
	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("error fetching user info: %w", err)
	}

	feed, err := db.GetFeed(context.Background(), url)
	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("error fetching feed info: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	createdFeedFollowRow, err := db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("error adding followed feed: %w", err)
	}

	return createdFeedFollowRow, nil
}
