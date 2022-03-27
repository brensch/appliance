package smarthome

import "fmt"

type ApplianceBase struct {
	X        int32
	Y        int32
	Health   int32
	Strength int32

	// TODO
	Repair int32
	Model  int32

	// The direction that Ability operates
	Pattern [8]bool
}

type Appliance interface {
	Ability()
	Attack()
}

type Toaster struct {
	ApplianceBase
}

func (t *Toaster) Attack() {
	fmt.Println(t.X, t.Y)
}

func (t *Toaster) Ability() {
	fmt.Println(t.Pattern)

}
