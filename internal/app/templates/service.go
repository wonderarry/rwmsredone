package templates

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/app"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type service struct{ provider app.TemplateProvider }

func New(provider app.TemplateProvider) Service { return &service{provider: provider} }

func (s *service) List(ctx context.Context) ([]domain.TemplateKey, error) {
	return s.provider.List(ctx)
}
func (s *service) Get(ctx context.Context, key domain.TemplateKey) (domain.CompiledTemplate, error) {
	return s.provider.Load(ctx, key)
}
