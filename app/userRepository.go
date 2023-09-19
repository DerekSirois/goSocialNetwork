package app

import "goSocialNetwork/models"

func (a *App) CreateUser(u *models.User) error {
	result := a.DB.Create(u)
	return result.Error
}

func (a *App) GetUserByUsername(username string) (*models.User, error) {
	u := &models.User{}
	result := a.DB.First(u, "username = ?", username)
	return u, result.Error
}
