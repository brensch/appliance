package smarthome

// Upgrade is something that can be added to appliances to give them boosts etc
type Upgrade interface {
	ReceiveEvents(appliances []Appliance, events []Event, turn uint8) []Event
}
