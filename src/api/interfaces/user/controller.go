package user

import (
	"api/domain"
	"api/interfaces"
	"api/usecases"
	"strconv"

	userUsecase "api/usecases/user"
)

// A UserController belong to the interface layer.
type Controller struct {
	UserInteractor userUsecase.UserInteractor
	Logger         usecases.Logger
}

// NewUserController returns the resource of users.
func NewUserController(sqlHandler interfaces.SQLHandler, logger usecases.Logger) *Controller {
	return &Controller{
		UserInteractor: userUsecase.UserInteractor{
			UserRepository: &Repository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of users.
func (uc *Controller) Index(c interfaces.Context) {
	u := domain.User{}
	c.Bind(&u)
	users, err := uc.UserInteractor.Index()
	if err != nil {
		c.JSON(500, interfaces.NewError(err))
		return
	}
	c.JSON(200, users)
}

// Show return response which contain the specified resource of a user.
func (uc *Controller) Show(c interfaces.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.UserInteractor.Show(id)
	if err != nil {
		c.JSON(500, interfaces.NewError(err))
		return
	}
	c.JSON(200, user)
}
