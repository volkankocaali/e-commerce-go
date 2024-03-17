package seed

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2/log"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/database"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/models/traits"
	"github.com/volkankocaali/e-commorce-go/pkg/repository"
	"strings"
)

type CategoriesSeeder struct {
	Active bool
	Count  int
}

func (c CategoriesSeeder) GetName() string {
	return "CategoriesSeeder"
}

func (c CategoriesSeeder) IsActive() bool {
	return c.Active
}

func (c CategoriesSeeder) generate() models.Categories {
	createdBy := gofakeit.RandomInt([]int{1, 100})

	return models.Categories{
		ID:          0,
		Name:        capitalizeWords(gofakeit.ProductCategory()),
		Description: gofakeit.ProductDescription(),
		Icon:        gofakeit.ImageURL(100, 100),
		Image: traits.Image{
			ImagePath: gofakeit.ImageURL(900, 900),
		},
		Active: true,
		Timestampable: traits.Timestampable{
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
}

func (c CategoriesSeeder) Seed() error {
	cfg := config.LoadConfig()
	db, _ := database.NewMysqlDB(*cfg)

	for i := 0; i < c.Count; i++ {
		categories := c.generate()
		categoriesRepo := repository.NewCategoriesRepository(db)
		_, err := categoriesRepo.Create(categories)

		if err != nil {
			log.Errorf("%s values: %+v\n", c.GetName(), categories)
			return fmt.Errorf("DB Insert Error: %w", err)
		}
	}

	fmt.Println("CategoriesSeeder is successful")
	return nil
}

func capitalizeWords(input string) string {
	words := strings.Fields(input)

	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, " ")
}
