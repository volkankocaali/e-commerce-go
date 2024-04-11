package _interface

import (
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
)

type ProductUseCase interface {
	ListProduct(page int, perPage int, productId *string) ([]schema.ProductResponseSchema, error)
}
