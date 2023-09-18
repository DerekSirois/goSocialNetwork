package app

import "goSocialNetwork/models"

func (a *App) GetAllPost() (*[]models.Post, error) {
	var p []models.Post
	result := a.DB.Find(&p)
	return &p, result.Error
}

func (a *App) GetPostById(id uint) (*models.Post, error) {
	var p models.Post
	result := a.DB.First(&p, id)
	return &p, result.Error
}

func (a *App) CreatePost(p *models.Post) error {
	result := a.DB.Create(p)
	return result.Error
}

func (a *App) UpdatePost(newPost models.Post, id uint) error {
	p, err := a.GetPostById(id)
	if err != nil {
		return err
	}
	result := a.DB.Model(&p).Updates(newPost)
	return result.Error
}

func (a *App) DeletePost(id uint) error {
	result := a.DB.Delete(&models.Post{}, id)
	return result.Error
}

func (a *App) SoftDeletePost(id uint) error {
	p, err := a.GetPostById(id)
	if err != nil {
		return err
	}
	result := a.DB.Model(&p).Updates(models.Post{Active: false})
	return result.Error
}
