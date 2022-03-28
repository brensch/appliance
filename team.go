package smarthome

type Team struct {
	Appliances []Appliance
}

// func ApplyTeamDelta(team []Appliance, delta )

// func (t Team) CreateEvents(g GameState) [2][]Event {

// 	var allEvents [2][]Event
// 	// var aggDeltas [2]MoveDelta
// 	for _, appliance := range t.Appliances {
// 		events := appliance.CreateEvents(g)
// 		allEvents[FriendsIndex] = append(allEvents[FriendsIndex], events[FriendsIndex]...)
// 		allEvents[EnemiesIndex] = append(allEvents[EnemiesIndex], events[EnemiesIndex]...)
// 	}

// 	return allEvents

// }
