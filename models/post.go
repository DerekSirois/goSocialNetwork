package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content    string
	Active     bool
	UserId     uint
	ResponseTo uint
}
