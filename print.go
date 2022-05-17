package smarthome

import (
	"fmt"
)

type GameSquare struct {
	Appliance Appliance
	Events    []Event
}

type GameSquare2 struct {
	Appliance *Appliance
	Events    []Event
}

const (
	SquareWidthBytes  = 15
	SquareHeightLines = 7
)

func PrintState2(width, height int8, state State2) {

	// fmt.Printf("%+v", appliances)
	board := make([]GameSquare2, width*height)
	// var houses [2]HouseState

	for _, appliance := range state.Appliances {
		// house, ok := appliance.(HouseState)
		// if ok {
		// 	if house.Team == -1 {
		// 		houses[0] = house
		// 	} else {
		// 		houses[1] = house
		// 	}
		// 	continue
		// }

		// board[GetIndex(width, appliance.State().Location.X, appliance.State().Location.Y)].Appliance = appliance
		board[GetIndex(width, appliance.Location.X, appliance.Location.Y)].Appliance = &appliance
	}

	for _, event := range state.Events {
		if event.Targets[0].Y < 0 || event.Targets[0].Y > 5 {
			continue
		}
		board[GetIndex(width, event.Targets[0].X, event.Targets[0].Y)].Events = append(board[GetIndex(width, event.Targets[0].X, event.Targets[0].Y)].Events, event)
	}

	canvas := make([][]string, width)
	for i := 0; i < int(width); i++ {
		canvas[i] = make([]string, height*SquareHeightLines)
	}

	for y := int8(0); y < height; y++ {
		for x := int8(0); x < width; x++ {
			square := board[GetIndex(width, x, y)]
			square.Print(x, y, canvas)
		}
	}

	// when printed, positive in y axis goes down
	// teamindex 0 == team 1, ie y+ => going down. their house therefore is at the top.
	// fmt.Printf("---------------  house t:%d h:%d  ---------------\n", houses[1].Team, houses[1].Health)
	PrintCanvas(canvas)
	// fmt.Printf("---------------  house t:%d h:%d  ---------------\n", houses[0].Team, houses[0].Health)

}

// puts all the rows for one square into the right string in the canvas
func (g GameSquare2) Print(x, y int8, canvas [][]string) {

	yBase := y * SquareHeightLines
	canvas[x][yBase] = "---------------"
	canvas[x][yBase+SquareHeightLines-1] = "---------------"
	if g.Appliance == nil {
		canvas[x][yBase+1] = fmt.Sprintf("empty")
		return
	}
	canvas[x][yBase+1] = fmt.Sprintf("(%d,%d) %s ", g.Appliance.Location.X, g.Appliance.Location.Y, g.Appliance.Type)
	canvas[x][yBase+2] = fmt.Sprintf("h:%d s:%d t:%d", g.Appliance.Health, g.Appliance.Strength, g.Appliance.Team)

	// eventBuf := bytes.NewBuffer(nil)
	for i, event := range g.Events {
		if i > SquareHeightLines-3 {
			fmt.Println("got too many events at once", len(g.Events))
			break
		}
		switch event.Type {
		case EventTypeModifyHealth:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("(%d,%d)[h%+d]", event.CausedBy.Location.X, event.CausedBy.Location.Y, event.Value)
		case EventTypeRelocate:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("(%d,%d)[r(%d,%d)]", event.CausedBy.Location.X, event.CausedBy.Location.Y, event.Targets[0].X, event.Targets[0].Y)
		case EventTypeApplianceDeath:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("(%d,%d)ded", event.CausedBy.Location.X, event.CausedBy.Location.Y)
		case EventTypeTurnStart:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("start turn %d", event.Turn)
		case EventTypeTurnEnd:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("end turn %d", event.Turn)

		}
	}

}

func PrintCanvas(canvas [][]string) {
	width := len(canvas)
	height := len(canvas[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%-15s ", canvas[x][y])
		}
		fmt.Println()
	}
}

func GetIndex(width, x, y int8) int8 {
	return x + width*y
}
