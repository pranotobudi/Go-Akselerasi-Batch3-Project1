package entity

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	UserID uint
	Name   string
}

type Movie struct {
	gorm.Model
	UserID uint
	Title  string
	Year   int
	Rating int
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
