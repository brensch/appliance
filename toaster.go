package smarthome

import "fmt"

type Toaster struct {
	ApplianceState
	// The direction that Ability operates
	// Pattern [8]bool
}

func (t Toaster) Type() ApplianceType {
	return "toaster"
}

func (t Toaster) MoveToStreet(width, height, team int8) Appliance {
	t.Location = t.Location.MoveToStreet(width, height, team)
	t.Team = team
	return t
}

// Toaster attacks if he's in the front row. He attacks straight.
// If he's not in the front row he attempts to move forward one row.
func (t Toaster) CreateEvents(appliances []Appliance) []Event {

	locationToAttack := Location{
		X: t.Location.X,
		Y: t.Location.Y + t.Team,
	}

	// if house in striking range, send.
	if LocationIsHouse(6, t.Team, locationToAttack) {
		return []Event{
			ModifyHouseHealthEvent{
				EventBase: EventBase{
					// Iteration: 0,
					CausedBy: t,
				},
				Team:  -t.Team,
				Value: -t.Strength,
			},
		}
	}

	// don't attack the edge of the map or fall off
	if !LocationValid(3, 6, locationToAttack) {
		return nil
	}

	for _, appliance := range appliances {
		// if this appliance is not in front of us, ignore it
		if !SameLocation(locationToAttack, appliance.State().Location) {
			continue
		}

		// attack if it's the other team in front of us
		if appliance.State().Team != t.Team {
			return []Event{
				ModifyHealthEvent{
					EventBase: EventBase{
						Iteration: 0,
						CausedBy:  t,
						Target:    locationToAttack,
					},
					Value: -t.Strength,
				},
			}
		}

		// do nothing if it's our own team in front of us
		fmt.Println("own team in front", t.Location, locationToAttack)
		return nil

	}

	return []Event{
		RelocationEvent{
			EventBase: EventBase{
				// Iteration: 0,
				CausedBy: t,
				Target:   t.Location,
			},
			NewLocation: locationToAttack,
		},
	}
}

func (t Toaster) ReceiveEvents(appliances []Appliance, events []Event, turn int8) (Appliance, []Event) {

	for _, event := range events {

		switch v := event.(type) {
		case ModifyHealthEvent:
			if SameLocation(v.EventBase.Target, t.Location) {
				t.Health += v.Value
			}
		case RelocationEvent:
			if SameLocation(v.EventBase.Target, t.Location) {
				t.Location = v.NewLocation

			}
		}
	}

	return t, nil
}
