package smarthome

import "fmt"

type Team struct {
	Appliances []Appliance
}

func DoMove(teams [2]Team) {
	fmt.Println(teams)
}
