package app

import (
	"fmt"

	"github.com/wonderarry/rwmsredone/internal/app/accounts"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/app/processes"
	"github.com/wonderarry/rwmsredone/internal/app/projects"
	"github.com/wonderarry/rwmsredone/internal/app/templates"
)

type Services struct {
	Accounts  accounts.Service
	Projects  projects.Service
	Processes processes.Service
	Templates templates.Service
}

type Deps struct {
	UoW       contract.UnitOfWork
	Templates contract.TemplateProvider
	IDGen     contract.IDGen

	PasswordHasher contract.PasswordHasher
	OIDCVerifier   contract.OIDCVerifier
	Clock          contract.Clock
}

func NewServices(d Deps) (Services, error) {
	if d.UoW == nil {
		return Services{}, fmt.Errorf("uow is nil")
	}

	if d.Templates == nil {
		return Services{}, fmt.Errorf("templates is nil")
	}

	if d.IDGen == nil {
		return Services{}, fmt.Errorf("idgen is nil")
	}

	return Services{
		accounts.New(d.UoW),
		projects.New(d.UoW, d.IDGen),
		processes.New(d.UoW, d.Templates, d.IDGen),
		templates.New(d.Templates),
	}, nil
}
