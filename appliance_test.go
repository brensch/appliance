package smarthome_test

import (
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
