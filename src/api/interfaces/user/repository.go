package user

import (
	"api/domain"
	"api/interfaces"
)

// A UserRepository belong to the interface layer
type Repository struct {
	SQLHandler interfaces.SQLHandler
}

// FindAll is returns the number of entities.
func (repo *Repository) FindAll() (users domain.Users, err error) {
	if err = repo.SQLHandler.Find(&users).Error; err != nil {
		return
	}
	return
}

// FindByID is returns the entity identified by the given id.
func (repo *Repository) FindByID(id int) (user domain.User, err error) {
	if err = repo.SQLHandler.Find(&user, id).Error; err != nil {
		return
	}
	return
}
