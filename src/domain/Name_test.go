package domain

import "testing"


func TestItReturnsValidName(t *testing.T) {
	name := NameEntity.New("Jorge")

	Testing.AssertEquals(t, "Jorge", name.Value(), "")
}