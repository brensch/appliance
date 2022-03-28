package smarthome_test

// type DoTeamMoveTestCase struct {
// 	Appliances []smarthome.Appliance
// 	MoveDeltas []smarthome.MoveDelta
// }

// var (
// 	doTeamMoveTests = []DoTeamMoveTestCase{
// 		{
// 			Appliances: []smarthome.Appliance{

// 				&smarthome.Toaster{
// 					ApplianceBase: smarthome.ApplianceBase{
// 						X:        0,
// 						Y:        1,
// 						Strength: 1,
// 					},
// 					Pattern: [8]bool{false, true, false, false, false, false, false, false},
// 				},
// 			},
// 		},
// 	}
// )

// func TestDoTeamMove(t *testing.T) {
// 	deltas := smarthome.DoTeamMove(doTeamMoveTests[0].Appliances)
// 	fmt.Println("deltas on enemy")
// 	// for _, move := range deltas[smarthome.EnemiesIndex] {
// 	// 	fmt.Println("normal damage")
// 	for _, delta := range deltas[smarthome.EnemiesIndex].NormalDamage {
// 		fmt.Println("x:", delta.X, "y:", delta.Y, "value:", delta.Value)
// 	}
// 	fmt.Println("special damage")
// 	for _, delta := range deltas[smarthome.EnemiesIndex].SpecialDamage {
// 		fmt.Println("x:", delta.X, "y:", delta.Y, "value:", delta.Value)
// 	}
// 	fmt.Println("healing")
// 	for _, delta := range deltas[smarthome.EnemiesIndex].Healing {
// 		fmt.Println("x:", delta.X, "y:", delta.Y, "value:", delta.Value)
// 	}
// 	fmt.Println("relocations")
// 	for _, delta := range deltas[smarthome.EnemiesIndex].Relocations {
// 		fmt.Println("start:", delta.StartX, delta.StartY, "end:", delta.EndX, delta.EndY)
// 	}
// 	// }

// 	fmt.Println("deltas on friends")
// 	// for _, move := range deltas[smarthome.FriendsIndex] {
// 	fmt.Println("normal damage")
// 	for _, delta := range deltas[smarthome.FriendsIndex].NormalDamage {
// 		fmt.Println("x:", delta.X, "y:", delta.Y, "value:", delta.Value)
// 	}
// 	fmt.Println("special damage")
// 	for _, delta := range deltas[smarthome.FriendsIndex].SpecialDamage {
// 		fmt.Println("x:", delta.X, "y:", delta.Y, "value:", delta.Value)
// 	}
// 	fmt.Println("healing")
// 	for _, delta := range deltas[smarthome.FriendsIndex].Healing {
// 		fmt.Println("x:", delta.X, "y:", delta.Y, "value:", delta.Value)
// 	}
// 	fmt.Println("relocations")
// 	for _, delta := range deltas[smarthome.FriendsIndex].Relocations {
// 		fmt.Println("start:", delta.StartX, delta.StartY, "end:", delta.EndX, delta.EndY)
// 	}
// 	// }

// }
