package _interface

import "github.com/volkankocaali/e-commorce-go/pkg/models"

type TagRepository interface {
	Create(tags models.Tags) (models.Tags, error)
	CreateTagProduct(tp models.ProductTags) (models.ProductTags, error)
}
