package smarthome

type Toaster struct {
	ObjectState
	// The direction that Ability operates
	// Pattern [8]bool
}

func (t Toaster) Type() ObjectType {
	return "toaster"
}

func (t Toaster) MoveToStreet(width, height, team int8) Appliance {
	t.Location = t.Location.MoveToStreet(width, height, team)
	t.Team = team
	return t
}

func (t Toaster) ReceiveEvents(appliances []Appliance, events []Event, turn uint8) ([]Appliance, []Event) {

	var newEvents []Event
	for _, event := range events {

		switch v := event.(type) {
		case ModifyHealthEvent:
			if SameLocation(v.Target, t.Location) {
				t.Health += v.Value
			}
		case RelocationEvent:
			if SameLocation(v.Target, t.Location) {
				t.Location = v.NewLocation
			}
		case TurnStartEvent:
			if SameLocation(v.Target, t.Location) {
				newEvents = append(newEvents, PushOrAttack(t, appliances)...)
			}
		}

		// if we're dead, return a nil appliance and emit a death event
		if t.Health <= 0 {
			return nil, append(newEvents,
				ApplianceDeathEvent{
					EventBase: event.Base(),
					Appliance: t,
				},
			)
		}
	}

	return []Appliance{t}, newEvents
}
