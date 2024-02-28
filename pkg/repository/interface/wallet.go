package _interface

type WalletRepository interface {
	CreateNewWallet(userID uint) (uint, error)
	ReferenceUserAddedPointsToWallet(referenceUser uint) error
}
