package main

type Square int

const (
	UNKNOWN Square = iota
	STAR
	NOTSTAR
)

func (sv Square) String() string {
	switch sv {
	case UNKNOWN:
		return " "
	case STAR:
		return "*"
	case NOTSTAR:
		return "x"
	}

	return " "
}
