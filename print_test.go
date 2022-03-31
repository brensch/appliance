package smarthome_test

import (
	"testing"

	"github.com/brensch/smarthome"
)

func TestPrintState(t *testing.T) {
	houseStates := [2]smarthome.HouseState{
		{
			Health:   3,
			Strength: 1,
			Team:     1,
		},
		{
			Health:   3,
			Strength: 1,
			Team:     -1,
		},
	}

	appliances := []smarthome.Appliance{
		// GoingUp
		smarthome.Toaster{
			ApplianceState: smarthome.ApplianceState{
				Team: 1,
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
				Team: -1,
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
				CausedBy: smarthome.Sticky{
					ApplianceState: smarthome.ApplianceState{
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
	state := smarthome.State{
		HouseStates: houseStates,
		Appliances:  appliances,
		Events:      events,
	}

	smarthome.PrintState(3, 6, state)

}
