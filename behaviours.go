package smarthome

// func PushOrAttack(appliance Appliance, appliances []Appliance) []Event {

// 	locationToAttack := Location{
// 		X: appliance.State().Location.X,
// 		Y: appliance.State().Location.Y + appliance.State().Team,
// 	}

// 	// // if house in striking range, send.
// 	// if LocationIsHouse(6, appliance.State().Team, locationToAttack) {
// 	// 	return []Event{
// 	// 		ModifyHouseHealthEvent{
// 	// 			EventBase: EventBase{
// 	// 				// Iteration: 0,
// 	// 				CausedBy: appliance,
// 	// 			},
// 	// 			Team:  -appliance.State().Team,
// 	// 			Value: -appliance.State().Strength,
// 	// 		},
// 	// 	}
// 	// }

// 	// // don't attack the edge of the map or fall off
// 	// if !LocationValid(3, 6, locationToAttack) {
// 	// 	return nil
// 	// }
// 	// if we're attacking a house, send it
// 	if locationToAttack.Y < 0 || locationToAttack.Y > 5 {
// 		return []Event{
// 			ModifyHealthEvent{
// 				EventBase: EventBase{
// 					Iteration: 0,
// 					CausedBy:  appliance,
// 					Target:    locationToAttack,
// 				},
// 				Value: -appliance.State().Strength,
// 			},
// 		}
// 	}

// 	var targetAppliance Appliance
// 	for _, otherAppliance := range appliances {
// 		// if this appliance is not in front of us, ignore it
// 		if SameLocation(locationToAttack, otherAppliance.State().Location) {
// 			targetAppliance = otherAppliance
// 			break
// 		}
// 	}

// 	// move up if there's no one to biff
// 	if targetAppliance == nil {

// 		return []Event{
// 			RelocationEvent{
// 				EventBase: EventBase{
// 					// Iteration: 0,
// 					CausedBy: appliance,
// 					Target:   appliance.State().Location,
// 				},
// 				NewLocation: locationToAttack,
// 			},
// 		}
// 	}

// 	// do nothing if there is someone but they're on the same team
// 	if appliance.State().Team == targetAppliance.State().Team {
// 		return nil
// 	}

// 	// otherwise, bifffff
// 	return []Event{
// 		ModifyHealthEvent{
// 			EventBase: EventBase{
// 				Iteration: 0,
// 				CausedBy:  appliance,
// 				Target:    locationToAttack,
// 			},
// 			Value: -appliance.State().Strength,
// 		},
// 	}

// }

// func StandAndAttack(appliance Appliance, appliances []Appliance) []Event {
// 	locationToAttack := Location{
// 		X: appliance.State().Location.X,
// 		Y: appliance.State().Location.Y + appliance.State().Team,
// 	}

// 	var targetAppliance Appliance
// 	for _, otherAppliance := range appliances {
// 		// if this appliance is not in front of us, ignore it
// 		if SameLocation(locationToAttack, otherAppliance.State().Location) {
// 			targetAppliance = otherAppliance
// 			break
// 		}
// 	}

// 	// do nothing if there's no target in front of us or the target is on the same team
// 	if targetAppliance == nil || appliance.State().Team == targetAppliance.State().Team {
// 		return nil
// 	}

// 	// otherwise, line em up
// 	return []Event{
// 		ModifyHealthEvent{
// 			EventBase: EventBase{
// 				CausedBy: appliance,
// 				Target:   locationToAttack,
// 			},
// 			Value: -appliance.State().Strength,
// 		},
// 	}

// }
