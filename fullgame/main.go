package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brensch/smarthome"
)

func main() {
	fmt.Println("starting game")

	appliances := []smarthome.Appliance{}

	fmt.Printf("you have %d appliances\n", len(appliances))

	fmt.Println("appliances available for selection")

	applianceOptions, upgradeOptions := smarthome.GenerateOptions(0)

	for i, applianceOption := range applianceOptions {
		fmt.Println(i, applianceOption.Type())
	}

	_ = upgradeOptions
	reader := bufio.NewReader(os.Stdin)

	var selectionIndex int
	var err error
	for {
		fmt.Printf("select appliance-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		selectionIndex, err = strconv.Atoi(text)
		if err != nil {
			panic("GAF")
		}

		if selectionIndex >= 0 && selectionIndex < len(applianceOptions) {
			break
		}
	}

	fmt.Println("got selection", applianceOptions[selectionIndex])

	var x int
	for {
		fmt.Printf("x-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		x, err = strconv.Atoi(text)
		if err != nil {
			panic("GAF")
		}

		if x > 0 && x < 3 {
			break
		}
	}

	var y int
	for {
		fmt.Printf("y-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		y, err = strconv.Atoi(text)
		if err != nil {
			panic("GAF")
		}

		if y > 0 && y < 3 {
			break
		}
	}

	selectedAppliance := applianceOptions[selectionIndex]

	selectedAppliance.SetLocation(smarthome.Location{
		X: int8(x),
		Y: int8(y),
	})

	fmt.Println(selectedAppliance.State().Location)

	appliances = append(appliances, selectedAppliance)
	fmt.Println(appliances)

	state := smarthome.State{
		Appliances: appliances,
	}
	smarthome.PrintState(3, 3, state)

}
