package store

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type SettingStore interface {
	Create(ctx context.Context, setting *v1.Setting) error
	Get(ctx context.Context, slug string) (*v1.Setting, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SettingList, error)
	Update(ctx context.Context, seting *v1.Setting) error
}
