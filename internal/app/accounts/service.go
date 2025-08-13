package accounts

import (
	"context"

	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type service struct {
	uow contract.UnitOfWork
}

func New(uow contract.UnitOfWork) Service {
	return &service{uow: uow}
}

func (s *service) GetMe(ctx context.Context, actorID domain.AccountID) (*Me, error) {
	var out *Me

	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		acc, err := tx.Accounts().Get(ctx, actorID)

		if err != nil {
			return err
		}
		roles, err := tx.Accounts().ListGlobalRoles(ctx, acc)

		if err != nil {
			return err
		}
		out = &Me{Account: *acc, Roles: roles}
		return nil
	})

	return out, err
}
