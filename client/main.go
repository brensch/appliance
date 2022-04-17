package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/brensch/smarthome"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	opt := option.WithCredentialsFile("./homebrawl-firebase-adminsdk-ai7ea-325c4c29eb.json")
	conf := &firebase.Config{ProjectID: "homebrawl"}

	ctx := context.Background()
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	gob.Register(smarthome.Toaster{})
	gob.Register(smarthome.Sticky{})

	snaps := client.Collection("test").Doc("send").Snapshots(context.Background())

	for {
		snap, err := snaps.Next()
		// DeadlineExceeded will be returned when ctx is cancelled.
		if status.Code(err) == codes.DeadlineExceeded {
			return
		}
		if err != nil {
			fmt.Printf("Snapshots.Next: %v\n", err)
			return
		}
		if !snap.Exists() {
			fmt.Printf("Document no longer exists\n")
			return
		}

		yeet, err := snap.DataAt("yeet")
		if err != nil {
			fmt.Println(err)
			return
		}

		reader := gob.NewDecoder(bytes.NewBuffer(yeet.([]byte)))
		var receivedAppliances []smarthome.Appliance
		err = reader.Decode(&receivedAppliances)
		if err != nil {
			log.Fatalf("Error on decode process: %v\n", err)
			return
		}
		for _, appliance := range receivedAppliances {
			switch v := appliance.(type) {
			case smarthome.Toaster:
				fmt.Printf("toaster: %d %d\n", v.Health, v.Strength)
			case smarthome.Sticky:
				fmt.Printf("sticky: %d %d\n", v.Health, v.Strength)
			}

		}
	}
}
