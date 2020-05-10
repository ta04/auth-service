package repository

import (
	authPB "github.com/ta04/auth-service/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	Store(*authPB.Auth1) (*authPB.C, error)
	ShowByUsername(*authPB.Auth2) (*authPB.Auth1, error)
}
