package interfaces

import "github.com/volkankocaali/e-commorce-go/pkg/schema"

type Helper interface {
	GenerateTokenClients(admin schema.UserSchemaResponse) (string, error)
	GeneratePasswordHash(password string) (string, error)
	GenerateReferralCode() (string, error)
	CompareHashAndPassword(a string, b string) error
}
