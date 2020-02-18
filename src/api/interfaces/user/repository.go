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
func (ur *Repository) FindAll() (users domain.Users, err error) {
	const query = `
		SELECT
			id,
			name
		FROM
			users
	`
	rows, err := ur.SQLHandler.Query(query)

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			return
		}
		user := domain.User{
			ID:   id,
			Name: name,
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

// FindByID is returns the entity identified by the given id.
func (ur *Repository) FindByID(userID int) (user domain.User, err error) {
	const query = `
		SELECT
			id,
			name
		FROM
			users
		WHERE
			id = ?
	`
	row, err := ur.SQLHandler.Query(query, userID)

	defer row.Close()

	if err != nil {
		return
	}

	var id int
	var name string
	row.Next()
	if err = row.Scan(&id, &name); err != nil {
		return
	}
	user.ID = id
	user.Name = name

	return
}
