package models

import (
	"time"
)

// AuthMstAdmin represents an admin user entity
type AuthMstAdmin struct {
	ID              uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Pubid           string     `json:"pubid" gorm:"uniqueIndex;size:50"`
	GroupID         uint64     `json:"group_id"`
	Nik             *string    `json:"nik" gorm:"size:20"`
	Name            string     `json:"name" gorm:"size:255"`
	Username        string     `json:"username" gorm:"size:255"`
	Email           string     `json:"email" gorm:"uniqueIndex;size:255"`
	Kontak          *string    `json:"kontak" gorm:"size:15"`
	Alamat          *string    `json:"alamat" gorm:"size:255"`
	Pict            *string    `json:"pict" gorm:"size:255"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"password" gorm:"size:255"`
	RememberToken   *string    `json:"remember_token" gorm:"size:100"`
	FlagFirstLogin  int8       `json:"flag_first_login" gorm:"default:1"`
	LastLogin       *time.Time `json:"last_login"`
	LastIP          *string    `json:"last_ip" gorm:"size:25"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	CreatedBy       *uint64    `json:"created_by"`
	UpdatedBy       *uint64    `json:"updated_by"`
	DeletedBy       *uint64    `json:"deleted_by"`
}

// TableName overrides the table name used by AuthMstAdmin to `auth_mst_admins`
func (AuthMstAdmin) TableName() string {
	return "auth_mst_admins"
}
