package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func sendMessage() {
	log.Println("starting firebase app")
	opts := option.WithCredentialsFile("aphrodite-fce41-firebase-adminsdk-fbsvc-88e7cb0256.json")
	config := &firebase.Config{ProjectID: "aphrodite-fce41"}
	app, err := firebase.NewApp(context.Background(), config, opts)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	log.Println("starting firebase messaging client")
	messagingClient, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("err creating Firebase Messaging client: %v", err)
	}

	log.Println("sending message")
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Test from sever",
			Body:  "This is a test of push notifications for go",
		},
		Token: "eYA23yAq1k41jZ46fjbbeQ:APA91bFNjW8XjY2ufPVJM4QuaRVkg_Ys7pMik465mnYGYZ5KxHmhijB0l6nIcPp6J7MU0yju06pzqJL45JPh6341aHLkX0GED00HkzTiJLmjI5jCAVJ_KWI",
	}

	response, err := messagingClient.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("error sending message: %v", err)
	}
	log.Println("message response", response)

}
