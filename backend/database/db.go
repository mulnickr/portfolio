package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func configureFirebase() (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "rmulnick-web"}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func ConfigureFirestore() (*firestore.Client, error) {
	app, err := configureFirebase()
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return client, nil
}
