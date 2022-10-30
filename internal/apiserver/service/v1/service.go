package v1

import "github.com/teamen/kays/internal/apiserver/store"

type Service interface {
	Users() UserSrv
	Categories() CategorySrv
	Settings() SettingSrv
	Orders() OrderSrv
	Customers() CustomerSrv
	Products() ProductSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{
		store,
	}
}
func (s *service) Products() ProductSrv {
	return newProducts(s)
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Categories() CategorySrv {
	return newCategories(s)
}
func (s *service) Settings() SettingSrv {
	return newSettings(s)
}

func (s *service) Orders() OrderSrv {
	return newOrders(s)
}

func (s *service) Customers() CustomerSrv {
	return newCustomers(s)
}
