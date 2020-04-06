package repository

import (
	authPB "github.com/SleepingNext/auth-service/proto"
)

type Repository interface{
	Store(*authPB.Auth1) (*authPB.C, error)
	ShowByUsername(*authPB.Auth2) (*authPB.Auth1, error)
}
