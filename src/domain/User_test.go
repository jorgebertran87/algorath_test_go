package domain

import "testing"

func TestItReturnsValidUser(t *testing.T) {
	id 	 := IdEntity.New("1")
	name := NameEntity.New("Jorge")

	user := UserEntity.New(IdEntity.New("1"), NameEntity.New("Jorge"))
	Testing.AssertEquals(t, id, user.Id(), "")
	Testing.AssertEquals(t, name, user.Name(), "")
}
