package seed

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gosimple/slug"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/database"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/models/traits"
	"github.com/volkankocaali/e-commorce-go/pkg/repository"
	"math"
	"math/rand"
	"strings"
)

const (
	LogMessage = "%s values: %+v\n"
)

type ProductSeeder struct {
	Active bool
	Count  int
}

func (p ProductSeeder) GetName() string {
	return "ProductSeeder"
}

func (p ProductSeeder) IsActive() bool {
	return p.Active
}

func (p ProductSeeder) generate() models.Product {
	product := gofakeit.Product()
	productName := strings.ToUpper(slug.Make(product.Name))
	categoriesName := strings.ToUpper(slug.Make(product.Categories[0]))
	sku := generateSKU(categoriesName[0:3], productName[0:5])
	price := product.Price

	discountedPrice := 0.00
	isDiscounted := gofakeit.Bool()
	if isDiscounted {
		discountPercentage := roundToDecimal(gofakeit.Float64Range(0.10, 0.40), 2)
		discountedPrice = calculateDiscountedPrice(price, discountPercentage)
	}
	shippingPrice := gofakeit.Float64Range(0.00, 39.00)

	return models.Product{
		SKU:         sku,
		ProductName: product.Name,
		Image: traits.Image{
			ImagePath: gofakeit.ImageURL(900, 900),
		},
		Size:          gofakeit.RandomString([]string{"S", "M", "L", "XL", "XXL"}),
		Stock:         gofakeit.Number(1, 100),
		RegularPrice:  price,
		DiscountPrice: roundToDecimal(discountedPrice, 2),
		ShippingPrice: roundToDecimal(shippingPrice, 2),
		Quantity:      gofakeit.Number(1, 100),
		ProductNote:   gofakeit.Quote(),
		CreatedBy:     1,
		UpdatedBy:     1,
		Status:        gofakeit.Bool(),
		Timestampable: traits.Timestampable{
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		ShortDescription:   product.Description[0:15],
		ProductDescription: product.Description,
	}
}

func (p ProductSeeder) Seed() error {
	cfg := config.LoadConfig()
	db, _ := database.NewMysqlDB(*cfg)
	productRepo := repository.NewProductRepository(db)
	tagRepo := repository.NewTagRepository(db)

	for i := 0; i < p.Count; i++ {
		product := p.generate()
		createProduct, err := productRepo.Create(product)

		if err != nil {
			log.Errorf(LogMessage, p.GetName(), product)
			return fmt.Errorf("product seeder db insert error: %w", err)
		}

		// Create product categories
		productCategories, err := productRepo.CreateProductCategories(models.ProductCategories{
			ProductId:    createProduct.ID,
			CategoriesId: uint(rand.Intn(100) + 1),
		})

		if err != nil {
			log.Errorf(LogMessage, p.GetName(), productCategories)
			return fmt.Errorf("product categories seeder db insert error: %w", err)
		}

		// Create product tags
		productTags, err := tagRepo.CreateTagProduct(models.ProductTags{
			ProductID: createProduct.ID,
			TagID:     uint(rand.Intn(100) + 1),
		})

		if err != nil {
			log.Errorf(LogMessage, p.GetName(), productTags)
			return fmt.Errorf("product tags db insert error: %w", err)
		}
	}

	fmt.Println("ProductSeeder is successful")
	return nil
}

func generateSKU(category, productCode string) string {
	return fmt.Sprintf("%s-%s-%d", category, productCode, rand.Intn(100000))
}

func calculateDiscountedPrice(price, discountPercentage float64) float64 {
	return price - (price * discountPercentage)
}

func roundToDecimal(x float64, decimalPlaces int) float64 {
	multiplier := math.Pow(10, float64(decimalPlaces))
	yuvarlanmisSayi := math.Round(x*multiplier) / multiplier
	return yuvarlanmisSayi
}
