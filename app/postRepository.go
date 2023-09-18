package app

import "goSocialNetwork/post"

func (a *App) GetAllPost() (*[]post.Post, error) {
	var p []post.Post
	result := a.DB.Find(&p)
	return &p, result.Error
}

func (a *App) GetPostById(id uint) (*post.Post, error) {
	var p post.Post
	result := a.DB.First(&p, id)
	return &p, result.Error
}

func (a *App) CreatePost(p *post.Post) error {
	result := a.DB.Create(p)
	return result.Error
}

func (a *App) UpdatePost(newPost post.Post, id uint) error {
	p, err := a.GetPostById(id)
	if err != nil {
		return err
	}
	result := a.DB.Model(&p).Updates(newPost)
	return result.Error
}

func (a *App) DeletePost(id uint) error {
	result := a.DB.Delete(&post.Post{}, id)
	return result.Error
}

func (a *App) SoftDeletePost(id uint) error {
	p, err := a.GetPostById(id)
	if err != nil {
		return err
	}
	result := a.DB.Model(&p).Updates(post.Post{Active: false})
	return result.Error
}
