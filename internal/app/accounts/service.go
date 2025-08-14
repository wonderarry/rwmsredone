package accounts

import (
	"context"
	"time"

	"github.com/wonderarry/rwmsredone/internal/app/contract"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type service struct {
	uow    contract.UnitOfWork
	hasher contract.PasswordHasher
	tokens contract.TokenIssuer
	idgen  contract.IDGen
}

func New(uow contract.UnitOfWork, hasher contract.PasswordHasher, tokens contract.TokenIssuer, idgen contract.IDGen) Service {
	return &service{uow: uow, hasher: hasher, tokens: tokens, idgen: idgen}
}

func (s *service) RegisterLocal(ctx context.Context, cmd RegisterLocal) (domain.AccountID, error) {
	if cmd.Login == "" || cmd.Password == "" {
		return "", domain.ErrInvalidCredentials
	}

	var accID domain.AccountID
	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		existing, err := tx.Identities().GetByProviderSubject(ctx, domain.ProviderLocal, cmd.Login)
		if err != nil {
			return err
		}
		if existing != nil {
			return domain.ErrAccAlreadyExists
		}

		accID = domain.AccountID(s.idgen.NewID())
		acc := &domain.Account{
			ID:          accID,
			FirstName:   cmd.FirstName,
			MiddleName:  cmd.MiddleName,
			LastName:    cmd.LastName,
			GroupNumber: cmd.Group,
		}
		if err := tx.Accounts().Create(ctx, acc); err != nil {
			return err
		}

		hash, err := s.hasher.Hash(cmd.Password)
		if err != nil {
			return err
		}

		// Not sure about the email for now
		email := ""
		idn := &domain.Identity{
			ID:           domain.IdentityID(s.idgen.NewID()),
			AccountID:    accID,
			Provider:     domain.ProviderLocal,
			Subject:      cmd.Login,
			Email:        &email,
			PasswordHash: &hash,
		}
		if err := tx.Identities().Create(ctx, idn); err != nil {
			return err
		}

		// TODO: Later when assigning global roles use this
		for _, r := range cmd.Roles {
			tx.Accounts().GrantGlobalRole(ctx, accID, r)
		}

		return tx.Outbox().Append(ctx, domain.AccountProvisioned{AccountID: accID, Provider: domain.ProviderLocal, Subject: cmd.Login})
	})
	return accID, err
}

func (s *service) LoginLocal(ctx context.Context, login, password string) (Token, error) {
	var out Token
	err := s.uow.WithTx(ctx, func(ctx context.Context, tx contract.Tx) error {
		idn, err := tx.Identities().GetByProviderSubject(ctx, domain.ProviderLocal, login)
		if err != nil {
			return err
		}
		if idn == nil || idn.PasswordHash == nil {
			return domain.ErrUnauthorized
		}
		if ok := s.hasher.Verify(*idn.PasswordHash, password); !ok {
			return domain.ErrUnauthorized
		}

		acc, err := tx.Accounts().Get(ctx, idn.AccountID)
		if err != nil {
			return err
		}

		roles, err := tx.Accounts().ListGlobalRoles(ctx, acc)

		if err != nil {
			return err
		}
		stringRoles := make([]string, 0, len(roles))
		for _, item := range roles {
			stringRoles = append(stringRoles, string(item))
		}

		// TODO: move to env
		tti := 24 * time.Hour

		jwt, err := s.tokens.Issue(ctx, string(acc.ID), map[string]interface{}{
			"aid":   acc.ID,
			"idp":   string(idn.Provider),
			"idsub": idn.Subject,
			"email": valueOrEmpty(idn.Email),
			"name":  acc.FirstName + " " + acc.LastName,
			"group": acc.GroupNumber,
			"roles": stringRoles,
			"scope": "user",
		}, tti)
		if err != nil {
			return err
		}
		out = Token{AccessToken: jwt}
		return nil
	})
	return out, err
}

/* ---------- Queries ---------- */

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

		ids, err := tx.Identities().ListByAccount(ctx, actorID)
		if err != nil {
			return err
		}
		infos := make([]IdentityInfo, 0, len(ids))
		for _, idn := range ids {
			info := ""
			if idn.Subject != "" {
				info = idn.Subject
			}
			email := ""
			if idn.Email != nil {
				email = *idn.Email
			}
			infos = append(infos, IdentityInfo{
				Provider: idn.Provider,
				Subject:  info,
				Email:    email,
			})
		}

		out = &Me{Account: *acc, Roles: roles, Identities: infos}
		return nil
	})
	return out, err
}

func valueOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
