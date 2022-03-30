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
						Team: 1,
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
						Team: 1,
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
						Team: -1,
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
	houseStates := [2]smarthome.HouseState{
		{
			Health:   3,
			Strength: 1,
			Team:     1,
		},
		{
			Health:   3,
			Strength: 1,
			Team:     -1,
		},
	}
	events := smarthome.CreateEvents(gameStateCreateEventsTests[0].Appliances)
	nextHouses, nextAppliances := smarthome.GetNextState(houseStates, gameStateCreateEventsTests[0].Appliances, events, 0)

	smarthome.PrintState(3, 6, nextHouses, nextAppliances, events)

	for _, appliance := range nextAppliances {
		fmt.Println(appliance.Type(), appliance.State().Location, appliance.State().Health)
	}

	// for i, team := range events {
	// 	for _, event := range team {
	// 		fmt.Println(i, event.Type())
	// 	}
	// }

}

var (
	houses = [2]smarthome.House{
		{
			Appliances: []smarthome.Appliance{
				// GoingUp
				smarthome.Toaster{
					ApplianceState: smarthome.ApplianceState{
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
						Location: smarthome.Location{
							X: 1,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
			},
			State: smarthome.HouseState{
				Health:   3,
				Strength: 3,
			},
		},
		{
			Appliances: []smarthome.Appliance{
				// GoingUp
				smarthome.Toaster{
					ApplianceState: smarthome.ApplianceState{
						Location: smarthome.Location{
							X: 0,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
				smarthome.Sticky{
					ApplianceState: smarthome.ApplianceState{
						Location: smarthome.Location{
							X: 1,
							Y: 2,
						},
						Strength: 1,
						Health:   3,
					},
				},
			},
			State: smarthome.HouseState{
				Health:   3,
				Strength: 3,
			},
		},
	}
)

func TestInitGame(t *testing.T) {

	result := smarthome.InitGame(houses)
	smarthome.PrintState(3, 6, result.HouseStates, result.Appliances, nil)
}

func TestGameResult(t *testing.T) {
	game := smarthome.InitGame(houses)
	result := game.Play()
	fmt.Println(result)
}

func TestPlay(t *testing.T) {
	game := smarthome.InitGame(houses)
	result := game.Play()
	fmt.Println("result", result)
}
