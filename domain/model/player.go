package model

type Player struct {
	ID string
	Name string
	SailNo int
	Sex string
	IconUrl string
	Point int
	Comment string
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}