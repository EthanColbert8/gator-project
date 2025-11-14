package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/EthanColbert8/gator-project/internal/database"
	"github.com/google/uuid"
)

func HandlerLogin(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 1)
	if err != nil {
		return err
	}

	_, err = s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error retrieving user: %w", err)
	}

	err = s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Printf("Current user set to: %s\n", cmd.Args[0])
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 1)
	if err != nil {
		return err
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	newUser, err := s.Db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	s.Cfg.SetUser(newUser.Name)
	fmt.Printf("User %s registered and set as current user.\n", newUser.Name)

	fmt.Printf("\nUser Details:\nID: %s\nCreated At: %s\nUpdated At: %s\n\n", newUser.ID, newUser.CreatedAt, newUser.UpdatedAt)
	return nil
}

func HandlerGetUsers(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 0)
	if err != nil {
		return err
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving users: %w", err)
	}

	var toPrint string
	for _, user := range users {
		if user.Name == s.Cfg.CurrentUsername {
			toPrint = fmt.Sprintf("* %s (current)", user.Name)
		} else {
			toPrint = fmt.Sprintf("* %s", user.Name)
		}

		fmt.Println(toPrint)
	}

	return nil
}

func HandlerResetUsers(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 0)
	if err != nil {
		return err
	}

	err = s.Db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting users table: %w", err)
	}

	return nil
}

func HandlerAggregate(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 0)
	if err != nil {
		return err
	}

	const feedUrl = "https://www.wagslane.dev/index.xml"
	feed, err := FetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Println(feed)
	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 2)
	if err != nil {
		return err
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	currentUser, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: feedName, Valid: true},
		Url:       feedUrl,
		UserID:    currentUser.ID,
	}

	_, err = s.Db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}

	fmt.Printf("Feed '%s' added for user '%s'.\n", feedName, currentUser.Name)

	_, err = FollowFeed(s.Db, s.Cfg.CurrentUsername, feedUrl)
	if err != nil {
		return fmt.Errorf("error following feed after adding: %w", err)
	}

	return nil
}

func HandlerListAllFeeds(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 0)
	if err != nil {
		return err
	}

	feeds, err := s.Db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s (%s) by %s\n", feed.Name.String, feed.Url, feed.UserName)
	}

	return nil
}

func HandlerFollowFeed(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 1)
	if err != nil {
		return err
	}

	returnedRow, err := FollowFeed(s.Db, s.Cfg.CurrentUsername, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	userName := returnedRow.UserName
	feedName := returnedRow.FeedName.String

	fmt.Printf("User \"%s\" is now following feed \"%s\".\n", userName, feedName)
	return nil
}

func HandlerGetFeedsForUser(s *State, cmd Command) error {
	err := validateNumArgs(cmd, 0)
	if err != nil {
		return err
	}

	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), s.Cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error fetching followed feeds: %w", err)
	}

	fmt.Printf("Feeds followed by %s:\n", s.Cfg.CurrentUsername)

	for _, follow := range feedFollows {
		fmt.Printf("* %s (%s)\n", follow.FeedName.String, follow.FeedUrl)
	}

	return nil
}
