package domain

type Id struct {
	id string
}

func (i Id) New(id string) Id {
	return Id{id}
}

func (i *Id) Id() string {
	return i.id
}


