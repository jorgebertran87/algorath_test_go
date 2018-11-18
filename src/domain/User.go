package domain

type User struct {
	id Id
	name Name
}

func (u User) New(id Id, name Name) User {
	return User{id, name}
}

func (u *User) Name() Name {
	return u.name
}

func (u *User) Id() Id {
	return u.id
}

