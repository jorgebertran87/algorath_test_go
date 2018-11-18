package domain

import "testing"


func TestItReturnsValidId(t *testing.T) {
	id := IdEntity.New("1")

	Testing.AssertEquals(t, "1", id.Value(), "")
}