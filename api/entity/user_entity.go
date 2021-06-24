package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role            string
	RolePermissions []RolePermission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Permission struct {
	gorm.Model
	Permission      string
	RolePermissions []RolePermission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type RolePermission struct {
	gorm.Model
	PermissionID string
	RoleID       string
}

type User struct {
	gorm.Model
	RoleID       uint
	Name         string
	Email        string
	Password     string
	Genres       []Genre       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Movies       []Movie       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MovieReviews []MovieReview `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
