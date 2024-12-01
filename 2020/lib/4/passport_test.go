package day4

import (
	"testing"
)

func TestValidHairColor(t *testing.T) {
	passport := NewPassport("hcl:#123abc")

	result := passport.validHairColor()

	if !result {
		panic("Unexpected")
	}
}

func TestInvalidHairColor(t *testing.T) {
	passport := NewPassport("hcl:#123abz")

	result := passport.validHairColor()

	if result {
		panic("Unexpected")
	}
}
