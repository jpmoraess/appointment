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

type RegisterOutput struct {
	TenantID int64
	UserID   int64
}

func (r *Register) Execute(ctx context.Context, input *RegisterInput) (output *RegisterOutput, err error) {
	result, err := r.store.RegisterTx(ctx, db.RegisterTxParams{
		Name:     input.Name,
		Password: input.Password,
	})

	if err != nil {
		return
	}

	output = &RegisterOutput{
		TenantID: result.TenantID,
		UserID:   result.UserID,
	}

	return
}
