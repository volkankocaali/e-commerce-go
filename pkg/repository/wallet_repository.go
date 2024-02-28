package repository

import (
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	interfaces "github.com/volkankocaali/e-commorce-go/pkg/repository/interface"
	"gorm.io/gorm"
)

type walletDatabase struct {
	DB *gorm.DB
}

func NewWalletRepository(DB *gorm.DB) interfaces.WalletRepository {
	return &walletDatabase{DB}
}

func (w *walletDatabase) CreateNewWallet(userID uint) (uint, error) {
	wallet := models.Wallet{
		UserId:  userID,
		Balance: 0,
	}
	err := w.DB.Create(&wallet).Error

	if err != nil {
		return 0, err
	}

	if err := w.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return 0, err
	}

	return wallet.ID, nil
}

func (w *walletDatabase) ReferenceUserAddedPointsToWallet(referenceUser uint) error {
	var wallet models.Wallet
	w.DB.Where("user_id = ?", referenceUser).First(&wallet)

	if wallet.ID == 0 {
		return nil
	}

	wallet.Balance += 20
	w.DB.Save(&wallet)

	return nil
}
