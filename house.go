package smarthome

type House struct {
	Appliances []Appliance
	State      HouseState
}

type HouseState struct {
	Health   int8
	Strength int8
	Team     int8
}

func (h HouseState) ReceiveDamage(events []Event) HouseState {
	for _, event := range events {
		e, ok := event.(ModifyHouseHealthEvent)
		if !ok {
			continue
		}

		if e.Team != h.Team {
			continue
		}

		h.Health += e.Value
	}

	return h
}
