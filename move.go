package smarthome

type MoveDelta struct {
	NormalDamage  []HealthModification
	SpecialDamage []HealthModification
	Healing       []HealthModification
	Relocations   []Relocation
}

type HealthModification struct {
	X     int32
	Y     int32
	Value int32
}

type Relocation struct {
	StartX int32
	StartY int32
	EndX   int32
	EndY   int32
}

// DoMove
func DoMove(appliances []Appliance) []MoveDelta {

	var allDeltas []MoveDelta
	for _, appliance := range appliances {
		moveDeltas := appliance.DoMove(appliances)
		allDeltas = append(allDeltas, moveDeltas)
	}
	return allDeltas

}
