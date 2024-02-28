package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

// TagParser interface defines the contract for tag parsing
type TagParser interface {
	Parse(pt []models.ProductTags) []schema.TagsResponseSchema
}
