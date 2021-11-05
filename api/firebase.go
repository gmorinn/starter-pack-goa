package api

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func firebaseClient(ctx context.Context) (*firebase.App, error) {
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	client, err := firebase.NewApp(ctx, config, option.WithCredentialsFile("../google.json"))
	if err != nil {
		return &firebase.App{}, fmt.Errorf("FIREBASE_NEWAPP %s", err.Error())
	}
	return client, nil
}
