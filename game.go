package smarthome

import "fmt"

// type Game struct {
// 	Appliances  []Appliance
// 	HouseStates [2]HouseState
// }

type State struct {
	Appliances []Appliance
	// HouseStates [2]HouseState
	Events []Event
}

type State2 struct {
	Appliances []Appliance
	// HouseStates [2]HouseState
	Events []Event
}

func GetFirstState(houses [2]House) State {

	var s State
	for i := 0; i < 2; i++ {
		// s.Appliances = []Appliance{houses[i].State}
		// s.HouseStates[i] = houses[i].State
		// s.HouseStates[i].Team = TeamValues[i]
		// move appliances into the street to fight
		// ie adjust their location if they are the second team
		for _, appliance := range houses[i].Appliances {
			streetAppliance := appliance.MoveToStreet(3, 6, TeamValues[i])
			s.Appliances = append(s.Appliances, streetAppliance)
		}
	}
	return s
}

func PlayGame(startingState State) ([]State, int8) {
	turn := uint8(0)

	states := []State{startingState}
	// PrintState(3, 6, houses, appliances, nil)

	for {
		if turn > 20 {
			fmt.Println("got to turn 20, get good please brend")
			return states, ResultTimeout
		}

		states = append(states, DoTurn(states[len(states)-1], turn)...)

		result := gameResult(states[len(states)-1])
		if result != ResultNotFinished {
			return states, result
		}

		turn++
	}
}

func TurnStartEvents(appliances []Appliance, turnNumber uint8) []Event {

	var allEvents []Event
	for _, appliance := range appliances {
		event := Event{
			Type:    EventTypeTurnStart,
			Targets: []Location{appliance.Location},
			Turn:    turnNumber,
		}
		allEvents = append(allEvents, event)
	}

	return allEvents
}

func TurnEndEvents(appliances []Appliance, turnNumber uint8) []Event {

	var allEvents []Event
	for _, appliance := range appliances {
		event := Event{
			Type:    EventTypeTurnEnd,
			Targets: []Location{appliance.Location},
			Turn:    turnNumber,
		}
		allEvents = append(allEvents, event)
	}

	return allEvents
}

func LoopUntilNoEventsRemaining(startingState State, turnNumber uint8) []State {
	var states []State
	prevState := startingState
	// states = append(states, prevState)
	loop := int8(0)

	// execute events until there are no new ones generated
	for len(prevState.Events) > 0 {
		var nextState State

		if loop > 100 {
			fmt.Println("got 100 iterations, get good please brend")
			return states
		}
		// // receive the damage from the previous events
		// // doing it here to make sure we get the events from the turn
		// for i, house := range prevState.HouseStates {
		// 	nextState.HouseStates[i] = house.ReceiveDamage(prevState.Events)
		// }

		// send each appliance all the previous events and see how they react
		for _, appliance := range prevState.Appliances {
			for _, event := range prevState.Events {

				nextAppliances, followUpEvents := appliance.ReceiveEvent(event, prevState.Appliances, turnNumber)
				nextState.Events = append(nextState.Events, followUpEvents...)
				nextState.Appliances = append(nextState.Appliances, nextAppliances...)
			}

		}

		states = append(states, nextState)
		prevState = nextState
		loop++

	}

	return states
}

func DoTurn(startingState State, turnNumber uint8) []State {
	// apply the deltas for each team.
	// Each team needs the events of the opposite team as their enemy, hence the [1-i]

	// generate all of the start of turn events
	startingStatePlusEvents := State{
		Events:     TurnStartEvents(startingState.Appliances, turnNumber),
		Appliances: startingState.Appliances,
		// HouseStates: startingState.HouseStates,
	}
	allStates := []State{startingStatePlusEvents}

	allStates = append(allStates, LoopUntilNoEventsRemaining(startingStatePlusEvents, turnNumber)...)

	// once all resultant events have finished, send the end of turn event and get all resultant states
	endingStatePlusEvents := State{
		Events:     TurnEndEvents(allStates[len(allStates)-1].Appliances, turnNumber),
		Appliances: allStates[len(allStates)-1].Appliances,
		// HouseStates: allStates[len(allStates)-1].HouseStates,
	}
	allStates = append(allStates, endingStatePlusEvents)
	allStates = append(allStates, LoopUntilNoEventsRemaining(allStates[len(allStates)-1], turnNumber)...)

	return allStates

}

const (
	ResultGoingDown   = int8(-1)
	ResultNotFinished = int8(0)
	ResultGoingUp     = int8(1)
	ResultDraw        = int8(2)
	ResultTimeout     = int8(3)
)

func gameResult(state State) int8 {

	if len(state.Appliances) == 0 {
		return ResultDraw
	}

	// Count the survivors from each type
	var score int8
	for _, appliance := range state.Appliances {
		score += appliance.Team
	}

	if score < 0 {
		return ResultGoingDown
	}

	if score > 0 {
		return ResultGoingUp
	}

	// If both sides have survivors, game isn't finished yet
	return ResultNotFinished

}
