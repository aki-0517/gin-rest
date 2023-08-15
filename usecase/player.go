package usecase

import (
	"gin-rest/domain/repository"
	"gin-rest/domain/model"
)

type Player interface {
	Create (name string) error
	Find (id int) (*model.Player, error)
	FindAll() ([]*model.Player, error)
}

type player struct {
	playerRepository repository.Player
}

func NewPlayer(r repository.Player) Player {
	return &player{r}
}

func (t *player) Create(name string) error {
	player := model.NewPlayer(name)
	if err := t.playerRepository.Create(player); err != nil {
		return err
	}
	return nil
}

func (t *player) Find(id int) (*model.Player, error) {
	player, err := t.playerRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (t *player) FindAll() ([]*model.Player, error) {
	player, err := t.playerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return player, nil
}