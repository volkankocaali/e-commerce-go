package database

import "github.com/volkankocaali/e-commorce-go/pkg/config"

type Database interface {
	Connect(config config.Config)
}
