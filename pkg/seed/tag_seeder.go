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
)

type TagSeeder struct {
	Active bool
	Count  int
}

func (t TagSeeder) GetName() string {
	return "TagSeeder"
}

func (t TagSeeder) IsActive() bool {
	return t.Active
}

func (t TagSeeder) generate() models.Tags {
	gofakeit.Seed(0)
	createdUser := gofakeit.RandomInt([]int{1, 2, 3, 4, 5})

	return models.Tags{
		TagName: gofakeit.Gamertag(),
		Icon:    gofakeit.ImageURL(100, 100),
		Timestampable: traits.Timestampable{
			CreatedAt: gofakeit.Date(),
			UpdatedAt: gofakeit.Date(),
		},
		CreatedBy: createdUser,
		UpdatedBy: createdUser,
	}
}

func (t TagSeeder) Seed() error {
	cfg := config.LoadConfig()
	db, _ := database.NewMysqlDB(*cfg)

	for i := 0; i < t.Count; i++ {
		tag := t.generate()
		tagRepo := repository.NewTagRepository(db)
		_, err := tagRepo.Create(tag)

		if err != nil {
			log.Errorf("%s values: %+v\n", t.GetName(), tag)
			return fmt.Errorf("DB Insert Error: %w", err)
		}
	}

	fmt.Println("TagSeeder is successful")
	return nil
}
