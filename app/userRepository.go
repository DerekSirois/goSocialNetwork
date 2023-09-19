package app

import "goSocialNetwork/models"

func (a *App) CreateUser(u *models.User) error {
	result := a.DB.Create(u)
	return result.Error
}
