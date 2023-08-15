package infla

import (
	"gin-rest/domain/repository"
	"gin-rest/domain/model"

	"gorm.io/gorm"
)

type Player struct {
	db *gorm.DB
}

func NewPlayer(db *gorm.DB) repository.Player {
	return &Player{
		db:db,
	}
}

func (pl *Player) Create(p *model.Player) error {
	if err := pl.db.Create(pl).Error; err != nil {
		return err
	}
	return nil
}

func (pl *Player) Find(id int) (*model.Player, error) {
	var player *model.Player
	err := pl.db.Where("id = ?", id).Take(&player).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return player, nil
}
func (pl *Player) FindAll() ([]*model.Player, error) {
	var players []*model.Player
	err := pl.db.Find(&players).Error
	if err != nil {
		return nil, err
	}
	return players, nil
}