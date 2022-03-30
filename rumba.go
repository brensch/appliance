package smarthome

// want him to path towards home
// currently just a copy of toaster
type Rumba struct {
	ApplianceState
}

func (t Rumba) Type() ApplianceType {
	return "rumba"
}

func (t Rumba) MoveToStreet(width, height, team int8) Appliance {
	t.Location = t.Location.MoveToStreet(width, height, team)
	t.Team = team
	return t
}

// Toaster attacks if he's in the front row. He attacks straight.
// If he's not in the front row he attempts to move forward one row.
func (t Rumba) CreateEvents(appliances []Appliance) []Event {

	locationToAttack := Location{
		X: t.Location.X,
		Y: t.Location.Y + t.Team,
	}

	for _, appliance := range appliances {
		// if this appliance is not in front of us, ignore it
		if !SameLocation(locationToAttack, appliance.State().Location) {
			continue
		}

		// do nothing if it's our own team in front of us
		if appliance.State().Team == t.Team {
			return nil
		}

		// attack if it's the other team in front of us
		return []Event{
			ModifyHealthEvent{
				EventBase: EventBase{
					CausedBy: t,
					Target:   locationToAttack,
				},
				Value: -t.Strength,
			},
		}

	}

	return []Event{
		RelocationEvent{
			EventBase: EventBase{
				CausedBy: t,
				Target:   t.Location,
			},
			NewLocation: locationToAttack,
		},
	}
}

func (t Rumba) ReceiveEvents(appliances []Appliance, events []Event, turn int8) (Appliance, []Event) {

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
