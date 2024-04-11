package seed

import (
	"github.com/gofiber/fiber/v2/log"
)

const DefaultActiveStatus = true

type Seeder interface {
	Seed() error
	GetName() string
	IsActive() bool
}

func CreateAndSeed() {
	userSeeder := UserSeeder{Active: DefaultActiveStatus, Count: 100}
	categorySeeder := CategoriesSeeder{Active: DefaultActiveStatus, Count: 300}
	tagSeeder := TagSeeder{Active: DefaultActiveStatus, Count: 100}
	productSeeder := ProductSeeder{Active: DefaultActiveStatus, Count: 200}
	run(userSeeder, categorySeeder, tagSeeder, productSeeder)
}

func run(seeders ...Seeder) {
	for _, seeder := range seeders {
		if seeder.IsActive() {
			if err := seeder.Seed(); err != nil {
				log.Fatalf("%s: %s", seeder.GetName(), err)
			}
		}
	}
}
