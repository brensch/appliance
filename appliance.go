package smarthome

// Positions x,y
// 0,2 1,2 2,2
// 0,1 1,1 2,1
// 0,0 1,0 2,0

type ApplianceState struct {
	Location Location
	// GoingUp is determined by the team the appliance is on
	GoingUp bool

	Health   int8
	Strength int8

	// TODO
	Repair int8
	Model  int8
}

type ApplianceType string

type Location struct {
	X int8
	Y int8
}

func SameLocation(loc1, loc2 Location) bool {
	return loc1.X == loc2.X && loc1.Y == loc2.Y
}

func (s ApplianceState) State() ApplianceState {
	return s
}

// const (
// 	FriendsIndex = 0
// 	EnemiesIndex = 1
// )

type Appliance interface {
	State() ApplianceState
	Type() ApplianceType
	// CreateEvents performs all the logic based on an appliances current position and the positions
	// of other appliances.
	// The friends and enemies inputs are two separate Appliance slices so that you can pass in different arrays
	// depending on the team of the appliance creating the events.
	CreateEvents([]Appliance) []Event
	// ReceiveEvent is how the appliance responds to events.
	// The returned appliance object is the next state of the appliance
	// All appliances receive all events in case they have behaviour based on things happening to
	// appliances elsewhere on the board.
	// Receiving events can also trigger more events, for example incoming damage to an ally might trigger a heal event etc
	ReceiveEvents([]Appliance, []Event) (Appliance, []Event)
}
