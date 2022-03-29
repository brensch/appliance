package smarthome

type EventType string

const (
	EventTypeModifyHealth = "modify_health"
	EventTypeRelocate     = "relocate"
)

type EventBase struct {
	Iteration int8
	// CausedBy is an appliance since team is important
	CausedBy Appliance
	Target   Location
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
