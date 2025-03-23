package models

import (
	"github.com/jinzhu/gorm"
)

type Manga struct {
	gorm.Model
	Title         string  `gorm:"not null"`
	AltTitle      *string
	URL           string  `gorm:"not null"`
	PublicURL     string  `gorm:"not null"`
	Rating        float32
	IsNSFW        bool
	CoverURL      string  `gorm:"not null"`
	LargeCoverURL *string
	State         *string
	Author        *string
	Source        string  `gorm:"not null"`
	Tags          []Tag   `gorm:"many2many:manga_tags;"`
}

func (m *Manga) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func GetMangaByID(db *gorm.DB, id uint) (*Manga, error) {
	var manga Manga
	if err := db.Preload("Tags").First(&manga, id).Error; err != nil {
		return nil, err
	}
	return &manga, nil
}

func ListManga(db *gorm.DB, offset, limit int) ([]Manga, error) {
	var manga []Manga
	if err := db.Preload("Tags").Offset(offset).Limit(limit).Find(&manga).Error; err != nil {
		return nil, err
	}
	return manga, nil
}

func (m *Manga) Update(db *gorm.DB) error {
	return db.Save(m).Error
}
