package smarthome

type EventType string

const (
	// EventTypeModifyHouseHealth = "modify_house_health"
	EventTypeModifyHealth    = "modify_health"
	EventTypeRelocate        = "relocate"
	EventTypeApplianceBirth  = "appliance_birth"
	EventTypeApplianceDeath  = "appliance_death"
	EventTypeApplianceBought = "appliance_bought"
	EventTypeApplianceSold   = "appliance_sold"
	EventTypeStartGame       = "start_game"
	EventTypeEndGame         = "end_game"
	EventTurnStart           = "turn_start"
	EventTurnEnd             = "turn_end"
	EventTypeHouseDeath      = "house_death"

	EventTypePlayerMovedAppliance  = "player_moved_appliance"
	EventTypePlayerBoughtAppliance = "player_bought_appliance"
	EventTypePlayerSoldAppliance   = "player_sold_appliance"
	EventTypePlayerBoughtItem      = "player_bought_item"
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
	// json.Unmarshaler
	Type() EventType
	Base() EventBase
}

type TurnStartEvent struct {
	EventBase
	Turn uint8
}

func (e TurnStartEvent) Type() EventType {
	return EventTurnStart
}

func (e TurnStartEvent) Base() EventBase {
	return e.EventBase
}

// func (e TurnStartEvent) MarshalJSON() ([]byte, error) {

// 	return e.EventBase
// }

type TurnEndEvent struct {
	EventBase
	Turn uint8
}

func (e TurnEndEvent) Type() EventType {
	return EventTurnEnd
}

func (e TurnEndEvent) Base() EventBase {
	return e.EventBase
}

type ModifyHealthEvent struct {
	EventBase
	Value int8
}

func (e ModifyHealthEvent) Type() EventType {
	return EventTypeModifyHealth
}

func (e ModifyHealthEvent) Base() EventBase {
	return e.EventBase
}

type RelocationEvent struct {
	EventBase
	NewLocation Location
}

func (e RelocationEvent) Type() EventType {
	return EventTypeRelocate
}

func (e RelocationEvent) Base() EventBase {
	return e.EventBase
}

type ApplianceBirthEvent struct {
	EventBase
	Appliance Appliance
}

func (e ApplianceBirthEvent) Type() EventType {
	return EventTypeApplianceBirth
}

func (e ApplianceBirthEvent) Base() EventBase {
	return e.EventBase
}

type ApplianceDeathEvent struct {
	EventBase
	Appliance Appliance
}

func (e ApplianceDeathEvent) Type() EventType {
	return EventTypeApplianceDeath
}

func (e ApplianceDeathEvent) Base() EventBase {
	return e.EventBase
}

type StartGameEvent struct {
	EventBase
}

func (e StartGameEvent) Type() EventType {
	return EventTypeStartGame
}

func (e StartGameEvent) Base() EventBase {
	return e.EventBase
}

type EndGameEvent struct {
	EventBase
}

func (e EndGameEvent) Type() EventType {
	return EventTypeEndGame
}

func (e EndGameEvent) Base() EventBase {
	return e.EventBase
}

// type ModifyHouseHealthEvent struct {
// 	EventBase
// 	Value int8
// 	Team  int8
// }

// func (e ModifyHouseHealthEvent) Type() EventType {
// 	return EventTypeModifyHouseHealth
// }

// // This kind of breaks the pattern since houses don't have a location per se.
// // Team will be all any event response requires.
// func (e ModifyHouseHealthEvent) Base() EventBase {
// 	return e.EventBase
// }

type BuyApplianceEvent struct {
	EventBase
	NewAppliance Appliance
}

func (e BuyApplianceEvent) Type() EventType {
	return EventTypeApplianceBought
}

func (e BuyApplianceEvent) Base() EventBase {
	return e.EventBase
}

type HouseDeathEvent struct {
	EventBase
}

func (e HouseDeathEvent) Type() EventType {
	return EventTypeHouseDeath
}

func (e HouseDeathEvent) Base() EventBase {
	return e.EventBase
}
