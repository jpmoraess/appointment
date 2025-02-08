package usecases

import (
	"context"

	db "github.com/jpmoraess/appointment-api/db/sqlc"
)

type Register struct {
	store db.Store
}

func NewRegister(store db.Store) *Register {
	return &Register{store: store}
}

type RegisterInput struct {
	Name     string
	Password string
}

func (r *Register) Execute(ctx context.Context, input *RegisterInput) error {
	_, err := r.store.RegisterTx(ctx, db.RegisterTxParams{
		Name:     input.Name,
		Password: input.Password,
	})

	if err != nil {
		return err
	}

	return nil
}
