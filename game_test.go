package smarthome_test

import (
	"fmt"
	"testing"

	"github.com/brensch/smarthome"
)

type GameStateCreateEventsCase struct {
	Appliances []smarthome.Appliance
	Events     []smarthome.Event
}

var (
	gameStateCreateEventsTests = []GameStateCreateEventsCase{
		{
			Appliances: []smarthome.Appliance{
				// GoingUp
				smarthome.Toaster{
					ApplianceState: smarthome.ApplianceState{
						GoingUp: true,
						Location: smarthome.Location{
							X: 0,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
				smarthome.Toaster{
					ApplianceState: smarthome.ApplianceState{
						GoingUp: true,
						Location: smarthome.Location{
							X: 1,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
				// GoingDown
				smarthome.Sticky{
					ApplianceState: smarthome.ApplianceState{
						GoingUp: false,
						Location: smarthome.Location{
							X: 0,
							Y: 3,
						},
						Strength: 1,
						Health:   3,
					},
				},
			},
		},
	}
)

func TestGameStateCreateEvents(t *testing.T) {
	events := smarthome.CreateEvents(gameStateCreateEventsTests[0].Appliances)

	for _, event := range events {
		fmt.Println(event.Type())
	}

	// for i, team := range events {
	// 	for _, event := range team {
	// 		fmt.Println(i, event.Type())
	// 	}
	// }

}

func TestGameStateGetNextState(t *testing.T) {
	events := smarthome.CreateEvents(gameStateCreateEventsTests[0].Appliances)
	nextAppliances := smarthome.GetNextState(gameStateCreateEventsTests[0].Appliances, events)

	smarthome.PrintState(3, 6, nextAppliances, events)

	for _, appliance := range nextAppliances {
		fmt.Println(appliance.Type(), appliance.State().Location, appliance.State().Health)
	}

	// for i, team := range events {
	// 	for _, event := range team {
	// 		fmt.Println(i, event.Type())
	// 	}
	// }

}

func TestPlay(t *testing.T) {
	result := smarthome.Play(gameStateCreateEventsTests[0].Appliances)
	fmt.Println(result)
}
