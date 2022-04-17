package smarthome

type Sticky struct {
	ObjectState
}

func (t Sticky) Type() ObjectType {
	return "sticky"
}

func (t Sticky) MoveToStreet(width, height, team int8) Appliance {
	t.Location = t.Location.MoveToStreet(width, height, team)
	t.Team = team
	return t
}

func (t Sticky) ReceiveEvents(appliances []Appliance, events []Event, turn uint8) ([]Appliance, []Event) {

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

		case TurnStartEvent:
			if SameLocation(v.EventBase.Target, t.Location) {
				newEvents = append(newEvents, StandAndAttack(t, appliances)...)
			}
		}

		// if we're dead, return a nil appliance and emit a death event
		if t.Health <= 0 {
			return nil, []Event{
				ApplianceDeathEvent{
					EventBase: event.Base(),
					Appliance: t,
				},
			}
		}

	}

	return []Appliance{t}, newEvents
}
