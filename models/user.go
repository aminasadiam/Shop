package models

import (
	"time"

	"gorm.io/gorm"
)

// Permission represents a single permission in the system
type Permission struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique"`
	Description string `gorm:"size:255"`
}

// Role represents a user role with associated permissions
type Role struct {
	gorm.Model
	Name        string       `gorm:"size:100;not null;unique"`
	Description string       `gorm:"size:255"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

// User represents a user in the system
type User struct {
	gorm.Model
	Email       string `gorm:"size:255;not null;unique"`
	Password    string `gorm:"size:255;not null"`
	FirstName   string `gorm:"size:100"`
	LastName    string `gorm:"size:100"`
	PhoneNumber string `gorm:"size:20"`
	IsActive    bool   `gorm:"default:true"`
	LastLoginAt time.Time
	Roles       []Role `gorm:"many2many:user_roles;"`
}

// TableName overrides the table name for User
func (User) TableName() string {
	return "users"
}

// TableName overrides the table name for Role
func (Role) TableName() string {
	return "roles"
}

// TableName overrides the table name for Permission
func (Permission) TableName() string {
	return "permissions"
}
