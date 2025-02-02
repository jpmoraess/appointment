package db

import "context"

type RegisterTxParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterTxResult struct {
	TenantID int64 `json:"tenantId"`
	UserID   int64 `json:"userId"`
}

func (store *SQLStore) RegisterTx(ctx context.Context, arg RegisterTxParams) (*RegisterTxResult, error) {
	var result *RegisterTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		tenant, err := q.CreateTenant(ctx, CreateTenantParams{
			Name: arg.Name,
			Slug: arg.Name,
		})
		if err != nil {
			return err
		}

		user, err := q.CreateUser(ctx, CreateUserParams{
			TenantID: tenant.ID,
			Username: arg.Name,
			Password: arg.Password,
			Role:     "admin",
		})
		if err != nil {
			return err
		}

		result = &RegisterTxResult{
			TenantID: tenant.ID,
			UserID:   user.ID,
		}

		return err
	})

	return result, err
}
