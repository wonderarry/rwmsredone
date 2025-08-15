package templates

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type service struct{ provider contract.TemplateProvider }

func New(provider contract.TemplateProvider) Service { return &service{provider: provider} }

func (s *service) List(ctx context.Context) ([]domain.TemplateKey, error) {
	return s.provider.List(ctx)
}
func (s *service) Get(ctx context.Context, key domain.TemplateKey) (domain.CompiledTemplate, error) {
	return s.provider.Load(ctx, key)
}
