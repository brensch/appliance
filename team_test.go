package smarthome_test

import (
	"fmt"
	"testing"

	"github.com/brensch/smarthome"
)

type DoMoveTestCase struct {
	Appliances []smarthome.Appliance
	MoveDeltas []smarthome.MoveDelta
}

var (
	doMoveTests = []DoMoveTestCase{
		{
			Appliances: []smarthome.Appliance{

				&smarthome.Toaster{
					ApplianceBase: smarthome.ApplianceBase{
						X: 0,
						Y: 0,
					},
				},
			},
		},
	}
)

func TestDoMove(t *testing.T) {
	deltas := smarthome.DoMove(doMoveTests[0].Appliances)
	fmt.Println(deltas)

}
