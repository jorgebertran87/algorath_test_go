package domain

type Id struct {
	id string
}

func (i Id) New(id string) Id {
	return Id{id}
}

func (i *Id) Value() string {
	return i.id
}


