package utils

import "testing"

func TestRound(t *testing.T) {
	if Round(2.1) != 2.0 {
		t.Fail()
	}

	if Round(2.6) != 3.0 {
		t.Fail()
	}
}

func TestRoundToPlaces(t *testing.T) {
	val := RoundToPlaces(2.34666, 2)
	if val != 2.35 {
		t.Errorf("val was %v", val)
	}
}
