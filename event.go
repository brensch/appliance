package smarthome

type EventType string

const (
	EventTypeModifyHealth      = "modify_health"
	EventTypeModifyHouseHealth = "modify_house_health"
	EventTypeRelocate          = "relocate"
	EventTypeCreateAppliance   = "create_appliance"
	EventTypeStartGame         = "start_game"
	EventTypeEndGame           = "end_game"
)

type EventBase struct {
	// Iteration is used for events that occur as a result of another event.
	// For instance the first time someone gets moved they might then get moved back as a result.
	// On the second move any appliance that responds to any move would then try to move them back again
	// creating an infinite loop.
	// By incrementing the iteration on the event we can check for iteration and only act a few times.
	Iteration int8
	// CausedBy is an appliance since team is important
	CausedBy Appliance
	// Target is only a location since it may not be specific to an appliance
	Target Location
}

type Event interface {
	// json.Marshaler
	Type() EventType
	Target() Location
}

type ModifyHealthEvent struct {
	EventBase
	Value int8
}

func (e ModifyHealthEvent) Type() EventType {
	return EventTypeModifyHealth
}

func (e ModifyHealthEvent) Target() Location {
	return e.EventBase.Target
}

type RelocationEvent struct {
	EventBase
	NewLocation Location
}

func (e RelocationEvent) Type() EventType {
	return EventTypeRelocate
}

func (e RelocationEvent) Target() Location {
	return e.EventBase.Target
}

type CreateApplianceEvent struct {
	EventBase
	Appliance Appliance
}

func (e CreateApplianceEvent) Type() EventType {
	return EventTypeCreateAppliance
}

func (e CreateApplianceEvent) Target() Location {
	return e.EventBase.Target
}

type StartGameEvent struct {
	EventBase
}

func (e StartGameEvent) Type() EventType {
	return EventTypeStartGame
}

func (e StartGameEvent) Target() Location {
	return e.EventBase.Target
}

type EndGameEvent struct {
	EventBase
}

func (e EndGameEvent) Type() EventType {
	return EventTypeEndGame
}

func (e EndGameEvent) Target() Location {
	return e.EventBase.Target
}

type ModifyHouseHealthEvent struct {
	EventBase
	Value int8
	Team  int8
}

func (e ModifyHouseHealthEvent) Type() EventType {
	return EventTypeModifyHouseHealth
}

// This kind of breaks the pattern since houses don't have a location per se.
// Team will be all any event response requires.
func (e ModifyHouseHealthEvent) Target() Location {
	return Location{}
}
