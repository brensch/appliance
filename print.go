package smarthome

import (
	"bytes"
	"fmt"
)

type GameSquare struct {
	Appliance Appliance
	Events    []Event
}

const (
	SquareWidthBytes  = 15
	SquareHeightLines = 5
)

func PrintState(width, height int8, appliances []Appliance, events []Event) {

	board := make([]GameSquare, width*height)

	for _, appliance := range appliances {
		board[GetIndex(width, appliance.State().Location.X, appliance.State().Location.Y)].Appliance = appliance
	}

	for _, event := range events {
		board[GetIndex(width, event.Target().X, event.Target().Y)].Events = append(board[GetIndex(width, event.Target().X, event.Target().Y)].Events, event)
	}

	// canvas := make([]byte, SquareWidthBytes*5*width*height)

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

	PrintCanvas(canvas)
}

func (g GameSquare) Print(x, y int8, canvas [][]string) {

	yBase := y * SquareHeightLines
	canvas[x][yBase] = "---------------"
	canvas[x][yBase+4] = "---------------"
	if g.Appliance == nil {
		canvas[x][yBase+1] = fmt.Sprintf("empty")

		return
	}
	canvas[x][yBase+1] = fmt.Sprintf("%s", g.Appliance.Type())
	canvas[x][yBase+2] = fmt.Sprintf("h:%d s:%d", g.Appliance.State().Health, g.Appliance.State().Strength)

	eventBuf := bytes.NewBuffer(nil)
	for _, event := range g.Events {
		switch v := event.(type) {
		case ModifyHealthEvent:
			eventBuf.WriteString(fmt.Sprintf("h%+d", v.Value))
		}
	}
	canvas[x][yBase+3] = eventBuf.String()
	// startLine := y*SquareHeightLines*width*SquareWidthBytes + xOffset
	// endLine := startLine + SquareWidthBytes
	// fmt.Println(canvas[startLine:endLine])
}

func PrintCanvas(canvas [][]string) {
	width := len(canvas)
	height := len(canvas[0])
	fmt.Println("heighto", height)
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
