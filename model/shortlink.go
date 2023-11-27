package model

import "gorm.io/gorm"

type Shortlink struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	LongLink string `json:"long_link" gorm:"uniqueIndex;not null"`
	TextRand string `json:"text_rand" gorm:"uniqueIndex; not null"`
	NameLink string `json:"name_link" gorm:"uniqueIndex;"`
	gorm.Model
}
