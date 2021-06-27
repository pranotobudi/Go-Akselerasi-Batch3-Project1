package entity

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	UserID uint
	Name   string
}

type Movie struct {
	gorm.Model
	UserID       uint
	Title        string
	Year         int
	Rating       int
	MovieReviews []MovieReview `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GenreMovies  []GenreMovie  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type GenreMovie struct {
	gorm.Model
	GenreID uint
	MovieID uint
}

type MovieReview struct {
	gorm.Model
	UserID  uint
	MovieID uint
	Review  string
	Rate    int
}
