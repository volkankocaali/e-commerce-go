package seed

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2/log"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/database"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"github.com/volkankocaali/e-commorce-go/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const UserPassword = "password"

type UserSeeder struct {
	Active bool
	Count  int
}

func (u UserSeeder) GetName() string {
	return "UserSeeder"
}

func (u UserSeeder) IsActive() bool {
	return u.Active
}

func (u UserSeeder) generate() models.Users {

	gofakeit.Seed(0)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Error("hash failed")
	}
	refCode, _ := generateReferralCode()

	return models.Users{
		Name:         gofakeit.Name(),
		Email:        gofakeit.Email(),
		Password:     string(hashedPassword),
		Phone:        gofakeit.Phone(),
		Blocked:      false,
		IsAdmin:      false,
		ReferralCode: refCode,
		BirthDate:    gofakeit.DateRange(time.Now().AddDate(-50, 0, 0), time.Now().AddDate(-30, 0, 0)).Format("2006-01-02"),
		Address: []models.Address{
			{
				Name:         gofakeit.RandomString([]string{"Home", "Work", "Other"}),
				Province:     gofakeit.Address().City,
				District:     gofakeit.Address().State,
				Neighborhood: gofakeit.Address().Street,
				FullAddress:  gofakeit.Address().Address,
				Phone:        gofakeit.Phone(),
				PostalCode:   gofakeit.Zip(),
				Country:      gofakeit.Address().Country,
				City:         gofakeit.Address().City,
				Default:      true,
			},
		},
	}
}

func (u UserSeeder) Seed() error {
	cfg := config.LoadConfig()
	db, _ := database.NewMysqlDB(*cfg)

	for i := 0; i < u.Count; i++ {
		user := u.generate()
		userRepo := repository.NewUserRepository(db)
		_, err := userRepo.Create(user)

		if err != nil {
			log.Errorf("%s values: %+v\n", u.GetName(), user)
			return fmt.Errorf("DB Insert Error: %w", err)
		}
	}

	fmt.Println("UserSeeder is successful")
	return nil
}

func generateReferralCode() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	encoded := base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding).EncodeToString(b)
	return encoded[:6], nil
}
