package smarthome

type ApplianceBase struct {
	Team     string
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
	DoMove([]Appliance) MoveDelta
}

type Toaster struct {
	ApplianceBase
}

func (t *Toaster) DoMove(appliances []Appliance) MoveDelta {
	return MoveDelta{}
}
