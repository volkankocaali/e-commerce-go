package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

// ProductParser interface defines the contract for product parsing
type ProductParser interface {
	Parse(pc []models.ProductCategories) ([]schema.ProductResponseSchema, error)
}
