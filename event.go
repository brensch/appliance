package smarthome

type EventType string

const (
	EventTypeModifyHealth = "modify_health"
	EventTypeRelocate     = "relocate"
)

type EventBase struct {
	Iteration int8
	CausedBy  Appliance
	Targets   Location
}

type Event interface {
	// json.Marshaler
	Type() EventType
}

type ModifyHealthEvent struct {
	EventBase
	Value int8
}

func (e ModifyHealthEvent) Type() EventType {
	return EventTypeModifyHealth
}

type RelocationEvent struct {
	EventBase
	NewLocation Location
}

func (e RelocationEvent) Type() EventType {
	return EventTypeRelocate
}
