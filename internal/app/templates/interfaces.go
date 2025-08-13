package templates

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Service interface {
	List(ctx context.Context) ([]domain.TemplateKey, error)
	Get(ctx context.Context, key domain.TemplateKey) (domain.CompiledTemplate, error)
}
