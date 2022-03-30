package smarthome_test

import (
	"testing"

	"github.com/brensch/smarthome"
)

func TestLocationValid(t *testing.T) {
	valid := smarthome.LocationValid(3, 6, smarthome.Location{X: 3, Y: 6})

	if valid {
		t.Log("shouldn't be valid")
		t.FailNow()
	}

	valid = smarthome.LocationValid(3, 6, smarthome.Location{X: 2, Y: 5})

	if !valid {
		t.Log("should be valid")
		t.FailNow()
	}

}

func TestLocationIsHouse(t *testing.T) {
	height := int8(6)
	isHouse := smarthome.LocationIsHouse(height, 1, smarthome.Location{X: 3, Y: 5})

	if isHouse {
		t.Log("shouldn't be house")
		t.FailNow()
	}

	isHouse = smarthome.LocationIsHouse(height, -1, smarthome.Location{X: 2, Y: -1})

	if !isHouse {
		t.Log("should be house")
		t.FailNow()
	}

}
