package helpers

import "testing"

func TestCreateRandomToken(t *testing.T) {
	result, err := CreateRandomToken(20)
	expect := "9x4694x79839g283x9g8"

	if result != expect {
		t.Errorf("Error: '%s', result:'%s', expected:'%s'", err, result, expect)
	}
}
