package smarthome

import "fmt"

// type GameState struct {
// 	Teams [2]Team
// }

func Play(appliances []Appliance) int {
	turns := 0

	for {
		if turns > 10 {
			fmt.Println("got 10 turns, get good please brend")
			return -2
		}
		turnEvents := CreateEvents(appliances)
		nextAppliances := GetNextState(appliances, turnEvents)
		result := gameResult(nextAppliances)

		if result != ResultNotFinished {
			return result
		}

		turns++
		appliances = nextAppliances
	}
}

func CreateEvents(appliances []Appliance) []Event {

	var allEvents []Event
	for _, appliance := range appliances {
		events := appliance.CreateEvents(appliances)
		// The index in allEvents that is the same as you will be your friends.
		allEvents = append(allEvents, events...)
		// The opposite index will be your enemies.
	}

	return allEvents

}

func GetNextState(appliances []Appliance, events []Event) []Appliance {
	// apply the deltas for each team.
	// Each team needs the events of the opposite team as their enemy, hence the [1-i]

	iterations := 0
	for len(events) > 0 {
		PrintState(3, 6, appliances, events)
		if iterations > 10 {
			fmt.Println("got 10 iterations, get good please brend")
			return appliances
		}
		fmt.Println("state iteration")
		var nextAppliances []Appliance
		var nextEvents []Event
		for _, appliance := range appliances {

			updatedAppliance, followUpEvents := appliance.ReceiveEvents(appliances, events)
			nextAppliances = append(nextAppliances, updatedAppliance)
			nextEvents = append(nextEvents, followUpEvents...)

		}

		events = nextEvents
		appliances = nextAppliances
		iterations++

	}

	// once there are no event left, remove all dead appliances.
	// this ensures that things like heal and whatnot works.
	var aliveAppliances []Appliance
	for _, appliance := range appliances {
		if appliance.State().Health < 0 {
			// TODO: make death event and iterate again
			continue
		}
		aliveAppliances = append(aliveAppliances, appliance)
	}

	// keeping as a value not a pointer for stack efficiency
	return aliveAppliances

}

const (
	ResultNotFinished = iota
	ResultDraw
	ResultGoingUp
	ResultGoingDown
)

func gameResult(appliances []Appliance) int {

	if len(appliances) == 0 {
		return ResultDraw
	}

	// Count the survivors from each type
	var goingUp, goingDown int8
	for _, appliance := range appliances {
		if appliance.State().GoingUp {
			goingUp++
			continue
		}
		goingDown++
	}

	if goingUp == 0 {
		return ResultGoingDown
	}

	if goingDown == 0 {
		return ResultGoingUp
	}

	// If both sides have survivors, game isn't finished yet
	return ResultNotFinished

}
