package smarthome_test

import (
	"fmt"
	"testing"

	"github.com/brensch/smarthome"
)

func TestApplySelection(t *testing.T) {

	appliances := []smarthome.Appliance{
		smarthome.HouseState{
			ObjectState: smarthome.ObjectState{
				Health:   3,
				Strength: 3,
			},
		},
		// GoingUp
		smarthome.Toaster{
			ObjectState: smarthome.ObjectState{
				Location: smarthome.Location{
					X: 0,
					Y: 2,
				},
				Strength: 1,
				Health:   3,
			},
		},
		smarthome.Sticky{
			ObjectState: smarthome.ObjectState{
				Location: smarthome.Location{
					X: 1,
					Y: 2,
				},
				Strength: 1,
				Health:   3,
			},
		},
	}

	selection := smarthome.Selection{
		Objects: appliances,
		PlayerEvent: smarthome.BuyApplianceEvent{
			NewAppliance: smarthome.Sticky{
				ObjectState: smarthome.ObjectState{
					Location: smarthome.Location{
						X: 2,
						Y: 2,
					},
					Strength: 1,
					Health:   3,
				},
			},
		},
	}

	newAppliances := smarthome.ApplySelection(selection)
	fmt.Println(len(newAppliances))

}
