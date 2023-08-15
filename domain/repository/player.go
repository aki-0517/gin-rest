package repository

import "gin-rest/domain/model"

type Player interface {
	Create(p *model.Player) error
	Find(id int) (*model.Player, error)
	FindAll() ([]*model.Player, error)
}