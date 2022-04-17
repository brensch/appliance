package smarthome

// the selection phase gives users an array of appliances and upgrades to choose from and
// the player generates a set of events to be applied to their game state

func GenerateOptions(turn int) ([]Appliance, []Upgrade) {

	appliances := []Appliance{
		// GoingUp
		Toaster{
			ObjectState: ObjectState{
				Location: Location{
					X: 0,
					Y: 2,
				},
				Strength: 1,
				Health:   3,
			},
		},
		Toaster{
			ObjectState: ObjectState{
				Location: Location{
					X: 1,
					Y: 2,
				},
				Strength: 1,
				Health:   3,
			},
		},
	}

	return appliances, nil
}

type Selection struct {
	Objects     []Appliance
	PlayerEvent Event
}

// players can rearrange their appliances without triggering events.
// only selection events that trigger other events are:
// - buy
// - sell
// - merge
// - upgrade
// for validation, will need a function that confirms all appliances listed in each event match previous known state,
// even if they've been rearranged
func ValidatePlayerSelection(previousState, nextState Selection) bool {
	// TODO: validation logic
	return true
}

func ApplySelection(selection Selection) []Appliance {

	var allNewAppliances []Appliance
	allEvents := []Event{selection.PlayerEvent}
	allAppliances := selection.Objects
	var allNewEvents []Event

	for len(allEvents) > 0 {
		for _, appliance := range allAppliances {
			newAppliances, newEvents := appliance.ReceiveEvents(allAppliances, allEvents, 0)
			allNewEvents = append(allNewEvents, newEvents...)
			allNewAppliances = append(allNewAppliances, newAppliances...)
		}
		allEvents = allNewEvents
		allAppliances = allNewAppliances
	}

	return allAppliances
}
