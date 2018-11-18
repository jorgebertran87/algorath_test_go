package domain

type Name struct {
	name string
}

func (n Name) New(name string) Name {
	return Name{name}
}

func (n *Name) Value() string {
	return n.name
}

