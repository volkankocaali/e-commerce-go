package repository

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (u *userDatabase) CheckUserExist(email string) bool {
	var user models.Users
	u.DB.Where("email = ?", email).First(&user)

	if user.ID != 0 {
		return true
	}

	return false
}

func (u *userDatabase) FindUserByReference(ref string) (uint, error) {
	var user models.Users
	u.DB.Where("referral_code = ?", ref).First(&user)

	if user.ID == 0 {
		return 0, nil
	}

	return user.ID, nil
}

func (u *userDatabase) FindUserByEmail(email string) (schema.UserSignInResponse, error) {
	var user models.Users
	u.DB.Where("email = ?", email).Preload("Address").First(&user)

	if user.ID == 0 {
		return schema.UserSignInResponse{}, nil
	}

	var address []schema.AddressSchemaResponse

	for _, v := range user.Address {
		address = append(address, schema.AddressSchemaResponse{
			Id:           v.Id,
			Name:         v.Name,
			Province:     v.Province,
			District:     v.District,
			Neighborhood: v.Neighborhood,
			FullAddress:  v.FullAddress,
			Phone:        v.Phone,
			PostalCode:   v.PostalCode,
			Country:      v.Country,
			City:         v.City,
			Default:      v.Default,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return schema.UserSignInResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		IsAdmin:   user.IsAdmin,
		BirthDate: user.BirthDate,
		Address:   address,
	}, nil
}

func (u *userDatabase) SignUp(user schema.UserSchema) (schema.UserSchemaResponse, error) {
	createdUser := models.Users{
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Phone:        user.Phone,
		ReferralCode: user.ReferralCode,
	}
	err := u.DB.Create(&createdUser).Error

	if err != nil {
		return schema.UserSchemaResponse{}, err
	}

	return schema.UserSchemaResponse{
		Id:    createdUser.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}

func (u *userDatabase) Create(user models.Users) (models.Users, error) {
	err := u.DB.Create(&user).Error

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}
func (u *userDatabase) CheckUserIsBlocked(email string) (bool, error) {
	var user models.Users
	u.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return false, nil
	}

	return user.Blocked, nil
}
