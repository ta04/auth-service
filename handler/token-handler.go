package handler

import userPB "github.com/ta04/user-service/proto"

// TokenHandler is the interface of token handlers.
// As there are number of token handlers can be used.
type TokenHandler interface {
	Encode(user *userPB.User) (string, error)
}
