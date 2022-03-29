package smarthome_test

import (
	"testing"

	"github.com/brensch/smarthome"
)

func TestPrintState(t *testing.T) {
	appliances := []smarthome.Appliance{
		// GoingUp
		smarthome.Toaster{
			ApplianceState: smarthome.ApplianceState{
				GoingUp: true,
				Location: smarthome.Location{
					X: 0,
					Y: 2,
				},
				Strength: 1,
				Health:   3,
			},
		},
		// GoingDown
		smarthome.Sticky{
			ApplianceState: smarthome.ApplianceState{
				GoingUp: false,
				Location: smarthome.Location{
					X: 0,
					Y: 3,
				},
				Strength: 1,
				Health:   3,
			},
		},
	}

	events := []smarthome.Event{
		smarthome.ModifyHealthEvent{
			EventBase: smarthome.EventBase{
				Iteration: 0,
				CausedBy: smarthome.Sticky{
					ApplianceState: smarthome.ApplianceState{
						GoingUp: false,
						Location: smarthome.Location{
							X: 0,
							Y: 3,
						},
						Strength: 1,
						Health:   3,
					},
				},
				Target: smarthome.Location{
					X: 0,
					Y: 2,
				},
			},
			Value: -1,
		},
	}

	smarthome.PrintState(3, 6, appliances, events)

}
