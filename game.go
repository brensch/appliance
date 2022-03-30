package smarthome

import "fmt"

type Game struct {
	Turn        uint8
	Appliances  []Appliance
	HouseStates [2]HouseState
}

func InitGame(houses [2]House) Game {

	var g Game
	for i := 0; i < 2; i++ {
		g.HouseStates[i] = houses[i].State
		g.HouseStates[i].Team = TeamValues[i]
		// move appliances into the street to fight
		// ie adjust their location if they are the second team
		for _, appliance := range houses[i].Appliances {
			streetAppliance := appliance.MoveToStreet(3, 6, TeamValues[i])
			g.Appliances = append(g.Appliances, streetAppliance)
		}
	}
	return g
}

func (g Game) Play() int8 {
	turn := int8(0)
	appliances := g.Appliances
	houses := g.HouseStates

	for {
		fmt.Printf("round %d-------------------------------------------------------\n", turn)
		if turn > 100 {
			fmt.Println("got 100 turn, get good please brend")
			return ResultTimeout
		}
		turnEvents := CreateEvents(appliances)
		nextHouses, nextAppliances := GetNextState(houses, appliances, turnEvents, turn)

		PrintState(3, 6, nextHouses, nextAppliances, nil)

		result := gameResult(nextHouses, nextAppliances)
		if result != ResultNotFinished {
			return result
		}

		turn++
		appliances = nextAppliances
		houses = nextHouses
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

func GetNextState(houses [2]HouseState, appliances []Appliance, events []Event, turn int8) ([2]HouseState, []Appliance) {
	// apply the deltas for each team.
	// Each team needs the events of the opposite team as their enemy, hence the [1-i]

	iteration := int8(0)
	// execute events until there are no new ones generated
	for len(events) > 0 {
		// fmt.Printf("turn iteration %d\n", iteration)
		// PrintState(3, 6, houses, appliances, events)
		if iteration > 100 {
			fmt.Println("got 100 iterations, get good please brend")
			return houses, appliances
		}
		// receive the damage from the previous events
		// doing it here to make sure we get the events from the turn
		var nextHouses [2]HouseState
		for i, house := range houses {
			nextHouses[i] = house.ReceiveDamage(events)
		}

		var nextAppliances []Appliance
		var nextEvents []Event
		for _, appliance := range appliances {
			updatedAppliance, followUpEvents := appliance.ReceiveEvents(appliances, events, turn)
			nextAppliances = append(nextAppliances, updatedAppliance)
			nextEvents = append(nextEvents, followUpEvents...)
		}

		events = nextEvents
		houses = nextHouses
		appliances = nextAppliances
		iteration++

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
	return houses, aliveAppliances

}

const (
	ResultGoingDown   = int8(-1)
	ResultNotFinished = int8(0)
	ResultGoingUp     = int8(1)
	ResultDraw        = int8(2)
	ResultTimeout     = int8(3)
)

func gameResult(houses [2]HouseState, appliances []Appliance) int8 {

	if len(appliances) == 0 {
		return ResultDraw
	}

	if houses[0].Health <= 0 && houses[1].Health <= 0 {
		return ResultDraw
	}

	for _, house := range houses {
		if houses[0].Health <= 0 {
			// return the opposite team to the one whos house has 0 health
			return -house.Team
		}
	}

	// Count the survivors from each type
	var score int8
	for _, appliance := range appliances {
		score += appliance.State().Team
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
