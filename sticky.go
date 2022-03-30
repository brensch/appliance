package smarthome

type Sticky struct {
	ApplianceState
}

func (t Sticky) Type() ApplianceType {
	return "sticky"
}

func (t Sticky) MoveToStreet(width, height, team int8) Appliance {
	t.Location = t.Location.MoveToStreet(width, height, team)
	t.Team = team
	return t
}

// Sticky doesn't let anyone on the other team move and attacks in the front row
func (t Sticky) CreateEvents(appliances []Appliance) []Event {

	locationToAttack := Location{
		X: t.Location.X,
		Y: t.Location.Y + t.Team,
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
						CausedBy: t,
						Target:   locationToAttack,
					},
					Value: -t.Strength,
				},
			}
		}

		// do nothing if it's our own team in front of us
		return nil
		// TODO: maybe make a relocate event and then tidy up using team tidy function
		// break
	}

	// do nothing if there's no-one in front of us
	return nil
}

func (t Sticky) ReceiveEvents(appliances []Appliance, events []Event, turn int8) (Appliance, []Event) {

	// Set yourself back where you were

	var newEvents []Event

	for _, event := range events {
		switch v := event.(type) {
		case RelocationEvent:
			if v.Iteration > 0 {
				continue
			}

			// ignore allies
			if v.CausedBy.State().Team == t.Team {
				continue
			}

			newEvents = append(newEvents,
				RelocationEvent{EventBase: EventBase{
					// increment the iteration so this event does not also get blocked
					Iteration: v.Iteration + 1,
					CausedBy:  t,
					Target:    v.NewLocation,
				},
					// use the original target as the new location since that's where they've come from
					NewLocation: v.EventBase.Target,
				})
		case ModifyHealthEvent:
			if SameLocation(v.EventBase.Target, t.Location) {
				t.Health += v.Value
			}
		}

	}

	return t, newEvents
}
