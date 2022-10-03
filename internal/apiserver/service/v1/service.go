package v1

import "github.com/teamen/kays/internal/apiserver/store"

type Service interface {
	Users() UserSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{
		store,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}
