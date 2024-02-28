package models

import "github.com/volkankocaali/e-commorce-go/pkg/models/traits"

type SlideShows struct {
	ID             int    `json:"id" gorm:"unique;not null"`
	DestinationUrl string `json:"destination_url"`
	ImageUrl       string `json:"image_url"`
	Clicks         int    `json:"clicks" gorm:"default:0"`
	traits.Timestampable
	CreatedBy int `json:"created_by"`
	UpdatedBy int `json:"updated_by"`
}
