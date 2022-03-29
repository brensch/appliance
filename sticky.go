package smarthome

import "fmt"

type Sticky struct {
	ApplianceState
}

func (t Sticky) Type() ApplianceType {
	return "sticky"
}

// Sticky doesn't let anyone on the other team move and attacks in the front row
func (t Sticky) CreateEvents(appliances []Appliance) []Event {

	targetY := t.Location.Y - 1
	if t.GoingUp {
		targetY += 2
	}
	locationToAttack := Location{
		X: t.Location.X,
		Y: targetY,
	}

	for _, appliance := range appliances {
		// if this appliance is not in front of us, ignore it
		if !SameLocation(locationToAttack, appliance.State().Location) {
			continue
		}

		// attack if it's the other team in front of us
		if appliance.State().GoingUp != t.GoingUp {
			return []Event{
				ModifyHealthEvent{
					EventBase: EventBase{
						Iteration: 0,
						CausedBy:  t,
						Target: Location{
							X: t.Location.X,
							Y: targetY,
						},
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

func (t Sticky) ReceiveEvents(appliances []Appliance, events []Event) (Appliance, []Event) {

	// Set yourself back where you were

	var newEvents []Event

	for _, event := range events {
		switch v := event.(type) {
		case RelocationEvent:
			if v.Iteration > 0 {
				continue
			}

			// ignore allies
			if v.CausedBy.State().GoingUp == t.GoingUp {
				continue
			}

			fmt.Printf("sticky saw an opponent trying to move from (%d,%d) to (%d,%d)\n", v.EventBase.Target.X, v.EventBase.Target.Y, v.NewLocation.X, v.NewLocation.Y)
			newEvents = append(newEvents,
				RelocationEvent{EventBase: EventBase{
					Iteration: v.Iteration + 1,
					CausedBy:  t,
					Target:    v.NewLocation,
				},
					// use the original target as the new location since that's where they've come from
					NewLocation: v.EventBase.Target,
				})
		case ModifyHealthEvent:
			if SameLocation(v.EventBase.Target, t.Location) {
				fmt.Println("sticky received a modify health event", v.Value)
				t.Health += v.Value
			}
		}

	}

	return t, newEvents
}
