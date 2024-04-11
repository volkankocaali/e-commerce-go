package models

import "github.com/volkankocaali/e-commorce-go/pkg/models/traits"

type Users struct {
	ID           uint      `json:"id" gorm:"unique;not null"`
	Name         string    `json:"name"`
	Email        string    `json:"email" validate:"email"`
	Password     string    `json:"password" validate:"min=8,max=20"`
	Phone        string    `json:"phone"`
	Blocked      bool      `json:"blocked" gorm:"default:false"`
	IsAdmin      bool      `json:"is_admin" gorm:"default:false"`
	ReferralCode string    `json:"referral_code"`
	BirthDate    string    `json:"birth_date"`
	Address      []Address `json:"addresses" gorm:"foreignKey:UserID"`
	traits.Timestampable
}

type Address struct {
	Id           uint   `json:"id" gorm:"unique;not null"`
	UserID       uint   `json:"user_id"`
	Users        Users  `json:"user" gorm:"foreignkey:UserID"`
	Name         string `json:"name" validate:"required"`
	Province     string `json:"province" validate:"required"`
	District     string `json:"district" validate:"required"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	FullAddress  string `json:"full_address" validate:"required"`
	Phone        string `json:"phone" gorm:"phone"`
	PostalCode   string `json:"postal_code" validate:"required"`
	Country      string `json:"country" validate:"required"`
	City         string `json:"city" validate:"required"`
	Default      bool   `json:"default" gorm:"default:false"`
	traits.Timestampable
}
