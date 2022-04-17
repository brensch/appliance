package smarthome_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"testing"

	"github.com/brensch/smarthome"
)

type MoveToStreetTestCase struct {
	StartLocation smarthome.Location
	EndLocation   smarthome.Location
	Width         int8
	Height        int8
	Team          int8
}

var (
	moveToStreetTestCases = []MoveToStreetTestCase{
		{
			StartLocation: smarthome.Location{X: 0, Y: 0},
			EndLocation:   smarthome.Location{X: 0, Y: 0},
			Width:         3,
			Height:        6,
			Team:          1,
		},
		{
			StartLocation: smarthome.Location{X: 0, Y: 0},
			EndLocation:   smarthome.Location{X: 2, Y: 5},
			Width:         3,
			Height:        6,
			Team:          -1,
		},
		{
			StartLocation: smarthome.Location{X: 2, Y: 0},
			EndLocation:   smarthome.Location{X: 0, Y: 5},
			Width:         3,
			Height:        6,
			Team:          -1,
		},
		{
			StartLocation: smarthome.Location{X: 0, Y: 5},
			EndLocation:   smarthome.Location{X: 2, Y: 0},
			Width:         3,
			Height:        6,
			Team:          -1,
		},
	}
)

func TestMoveToStreet(t *testing.T) {

	for _, test := range moveToStreetTestCases {
		calculatedTestLocation := test.StartLocation.MoveToStreet(test.Width, test.Height, test.Team)
		if !smarthome.SameLocation(test.EndLocation, calculatedTestLocation) {
			t.Logf("got %+v, expecting %+v", calculatedTestLocation, test.EndLocation)
			t.FailNow()
		}
	}
}

func TestInterfaces(t *testing.T) {
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
			Health:   3,
		},
	}

	appliances := []smarthome.Appliance{toaster, sticky, toaster}

	gob.Register(smarthome.Toaster{})
	gob.Register(sticky)

	var buf bytes.Buffer
	writer := gob.NewEncoder(&buf)
	err := writer.Encode(appliances)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	reader := gob.NewDecoder(&buf)
	var receivedAppliances []smarthome.Appliance
	err = reader.Decode(&receivedAppliances)
	if err != nil {
		log.Fatalf("Error on decode process: %v\n", err)
		return
	}

	for _, appliance := range appliances {

		fmt.Println(appliance.Type())
	}
}
