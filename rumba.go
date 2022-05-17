package smarthome

// // want him to path towards home
// // currently just a copy of toaster
// type Rumba struct {
// 	*ObjectState
// }

// func (t Rumba) Type() ObjectType {
// 	return "rumba"
// }

// func (t Rumba) MoveToStreet(width, height, team int8) Appliance {
// 	t.Location = t.Location.MoveToStreet(width, height, team)
// 	t.Team = team
// 	return t
// }

// func (t Rumba) ReceiveEvents(appliances []Appliance, events []Event, turn uint8) ([]Appliance, []Event) {

// 	var newEvents []Event
// 	for _, event := range events {

// 		switch v := event.(type) {
// 		case ModifyHealthEvent:
// 			if SameLocation(v.EventBase.Target, t.Location) {
// 				t.Health += v.Value
// 			}
// 		case RelocationEvent:
// 			if SameLocation(v.EventBase.Target, t.Location) {
// 				t.Location = v.NewLocation
// 			}
// 		case TurnStartEvent:
// 			if SameLocation(v.EventBase.Target, t.Location) {
// 				newEvents = append(newEvents, PushOrAttack(t, appliances)...)
// 			}
// 		}

// 		// if we're dead, return a nil appliance and emit a death event
// 		if t.Health <= 0 {
// 			return nil, append(newEvents,
// 				ApplianceDeathEvent{
// 					EventBase: event.Base(),
// 					Appliance: t,
// 				},
// 			)
// 		}
// 	}

// 	return []Appliance{t}, newEvents
// }
