package smarthome

import (
	"fmt"
)

type GameSquare struct {
	Appliance Appliance
	Events    []Event
}

const (
	SquareWidthBytes  = 15
	SquareHeightLines = 7
)

func PrintState(width, height int8, state State) {

	// fmt.Printf("%+v", appliances)
	board := make([]GameSquare, width*height)
	var houses [2]HouseState

	for _, appliance := range state.Appliances {
		house, ok := appliance.(HouseState)
		if ok {
			if house.Team == -1 {
				houses[0] = house
			} else {
				houses[1] = house
			}
			continue
		}

		board[GetIndex(width, appliance.State().Location.X, appliance.State().Location.Y)].Appliance = appliance
	}

	for _, event := range state.Events {
		if event.Base().Target.Y < 0 || event.Base().Target.Y > 5 {
			continue
		}
		board[GetIndex(width, event.Base().Target.X, event.Base().Target.Y)].Events = append(board[GetIndex(width, event.Base().Target.X, event.Base().Target.Y)].Events, event)
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
func (g GameSquare) Print(x, y int8, canvas [][]string) {

	yBase := y * SquareHeightLines
	canvas[x][yBase] = "---------------"
	canvas[x][yBase+SquareHeightLines-1] = "---------------"
	if g.Appliance == nil {
		canvas[x][yBase+1] = fmt.Sprintf("empty")
		return
	}
	canvas[x][yBase+1] = fmt.Sprintf("(%d,%d) %s ", g.Appliance.State().Location.X, g.Appliance.State().Location.Y, g.Appliance.Type())
	canvas[x][yBase+2] = fmt.Sprintf("h:%d s:%d t:%d", g.Appliance.State().Health, g.Appliance.State().Strength, g.Appliance.State().Team)

	// eventBuf := bytes.NewBuffer(nil)
	for i, event := range g.Events {
		if i > SquareHeightLines-3 {
			fmt.Println("got too many events at once", len(g.Events))
			break
		}
		switch v := event.(type) {
		case ModifyHealthEvent:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("(%d,%d)[h%+d]", v.CausedBy.State().Location.X, v.CausedBy.State().Location.Y, v.Value)
		case RelocationEvent:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("(%d,%d)[r(%d,%d)]", v.CausedBy.State().Location.X, v.CausedBy.State().Location.Y, v.NewLocation.X, v.NewLocation.Y)
		case ApplianceDeathEvent:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("(%d,%d)ded", v.CausedBy.State().Location.X, v.CausedBy.State().Location.Y)
		case TurnStartEvent:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("start turn %d", v.Turn)
		case TurnEndEvent:
			canvas[x][yBase+3+int8(i)] = fmt.Sprintf("end turn %d", v.Turn)

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
