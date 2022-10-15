package v1

import (
	"context"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type SettingSrv interface {
	Create(ctx context.Context, setting *v1.Setting) error

	Get(ctx context.Context, slug string) (*v1.Setting, error)
}

var _ SettingSrv = (*settingService)(nil)

type settingService struct {
	store store.Factory
}

func newSettings(srv *service) *settingService {
	return &settingService{
		store: srv.store,
	}
}

func (s *settingService) Create(ctx context.Context, setting *v1.Setting) error {
	if err := s.store.Settings().Create(ctx, setting); err != nil {
		return err
	}
	return nil
}

func (s *settingService) Get(ctx context.Context, slug string) (*v1.Setting, error) {

	setting, err := s.store.Settings().Get(ctx, slug)
	if err != nil {
		return nil, err
	}
	return setting, nil
}
