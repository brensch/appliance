package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/brensch/smarthome"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {

	// Use the application default credentials
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

	toaster := smarthome.Toaster{
		ObjectState: smarthome.ObjectState{
			Team: 1,
			Location: smarthome.Location{
				X: 0,
				Y: 2,
			},
			Strength: 1,
			Health:   3,
		},
	}
	sticky := smarthome.Sticky{
		ObjectState: smarthome.ObjectState{
			Location: smarthome.Location{
				X: 1,
				Y: 2,
			},
			Strength: 1,
			Health:   69,
		},
	}
	appliances := []smarthome.Appliance{toaster, sticky, toaster}

	gob.Register(smarthome.Toaster{})
	gob.Register(smarthome.Sticky{})

	var buf bytes.Buffer
	writer := gob.NewEncoder(&buf)
	err = writer.Encode(appliances)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("setting")
	_, err = client.Collection("test").Doc("send").Set(context.Background(), map[string][]byte{
		"yeet": buf.Bytes(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// snap, err := client.Collection("test").Doc("send").Get(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// yeet, err := snap.DataAt("yeet")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(yeet)
	// reader := gob.NewDecoder(bytes.NewBuffer(yeet.([]byte)))
	// var receivedAppliances []smarthome.Appliance
	// err = reader.Decode(&receivedAppliances)
	// if err != nil {
	// 	log.Fatalf("Error on decode process: %v\n", err)
	// 	return
	// }
	// for _, appliance := range appliances {

	// 	fmt.Println(appliance.Type())
	// }
}
