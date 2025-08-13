package app

import (
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
