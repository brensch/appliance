package smarthome_test

import (
	"testing"

	"github.com/brensch/smarthome"
)

type DoMoveTestCase struct {
	TeamsBefore [2]smarthome.Team
	TeamsAfter  [2]smarthome.Team
}

var (
	doMoveTests = []DoMoveTestCase{
		{
			TeamsBefore: [2]smarthome.Team{
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
			},
		},
	}
)

func TestDoMove(t *testing.T) {

	smarthome.DoMove(doMoveTests[0].TeamsBefore)
	for _, team := range doMoveTests[0].TeamsBefore {
		for _, appliance := range team.Appliances {

			appliance.Attack()
		}
	}

}
