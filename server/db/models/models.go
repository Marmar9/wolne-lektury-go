package models

import "gorm.io/gorm"

type Lecture struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint
	Author   Author
	KindID   uint
	Kind     Kind
}

type Author struct {
	gorm.Model
	Name     string    `json:"name"`
	Lectures []Lecture `gorm:"foreignKey:AuthorID"` // One-to-Many relationship
}

type Kind struct {
	gorm.Model
	Name     string    `json:"name"`
	Lectures []Lecture `gorm:"foreignKey:KindID"` // One-to-Many relationship
}
