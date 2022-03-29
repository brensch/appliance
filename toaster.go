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

// Toaster attacks if he's in the front row. He attacks straight.
// If he's not in the front row he attempts to move forward one row.
func (t Toaster) CreateEvents(appliances []Appliance) []Event {

	targetY := t.Location.Y - 1
	if t.GoingUp {
		targetY += 2
	}
	locationToAttack := Location{
		X: t.Location.X,
		Y: targetY,
	}
	fmt.Println(locationToAttack)

	for _, appliance := range appliances {
		// if this appliance is not in front of us, ignore it
		if !SameLocation(locationToAttack, appliance.State().Location) {
			continue
		}

		// attack if it's the other team in front of us
		fmt.Println("toaster attacking appliance in front of it")
		if appliance.State().GoingUp != t.GoingUp {
			return []Event{
				ModifyHealthEvent{
					EventBase: EventBase{
						Iteration: 0,
						CausedBy:  t,
						Target:    appliance.State().Location,
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

	// move forward if we didn't find someone to attack and aren't blocked
	newY := t.Location.Y - 1
	if t.GoingUp {
		newY += 2
	}

	return []Event{
		RelocationEvent{
			EventBase: EventBase{
				Iteration: 0,
				CausedBy:  t,
				Target:    t.Location,
			},
			NewLocation: locationToAttack,
		},
	}

	// return moveDeltas
}

func (t Toaster) ReceiveEvents(appliances []Appliance, events []Event) (Appliance, []Event) {

	for _, event := range events {

		switch v := event.(type) {
		case ModifyHealthEvent:
			if SameLocation(v.EventBase.Target, t.Location) {
				fmt.Println("toaster received a modify health event", v.Value)
				t.Health += v.Value
			}
		case RelocationEvent:
			if SameLocation(v.EventBase.Target, t.Location) {
				fmt.Println("toaster received a relocation event")
				t.Location = v.NewLocation

			}
		}
	}

	return t, nil
}
