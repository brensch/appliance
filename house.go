package smarthome

type House struct {
	Appliances []Appliance
}

// HouseState implements Appliance since it can be responsible for events, and receive damage
type HouseState struct {
	// May rethink the appliancestate interface
	ObjectState
	// Health   int8
	// Strength int8
	// Team     int8
}

const ObjectTypeHouse = "house"

func (h HouseState) Type() ObjectType {
	return ObjectTypeHouse
}

func (h HouseState) ReceiveEvents(appliances []Appliance, events []Event, turn uint8) ([]Appliance, []Event) {

	var newEvents []Event
	var newAppliances []Appliance
	for _, event := range events {

		switch v := event.(type) {
		case ModifyHealthEvent:
			// TODO: figure out if i want dynamically sized boards. probs not
			// if team unset (ie 0) then also consider cases when y<0
			if (v.Target.Y > 5 && h.Team == -1) || (v.Target.Y < 0 && h.Team >= 0) {
				h.Health += v.Value
			}
		// The housestate is tasked with responding to the purchase of appliances
		case BuyApplianceEvent:
			newAppliances = append(newAppliances, v.NewAppliance)
		}

		// if we're dead, return a nil appliance and emit a death event
		if h.Health <= 0 {
			return nil, append(events,
				HouseDeathEvent{
					EventBase: event.Base(),
				},
			)
		}
	}

	return append(newAppliances, h), newEvents
}

func (h HouseState) MoveToStreet(width, height, team int8) Appliance {
	h.Location = h.Location.MoveToStreet(width, height, team)
	h.Team = team
	return h
}

// func (h HouseState) ReceiveDamage(events []Event) HouseState {
// 	for _, event := range events {
// 		e, ok := event.(ModifyHouseHealthEvent)
// 		if !ok {
// 			continue
// 		}

// 		if e.Team != h.Team {
// 			continue
// 		}

// 		h.Health += e.Value
// 	}

// 	return h
// }
