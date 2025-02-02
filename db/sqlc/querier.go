// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetTenant(ctx context.Context, id int64) (Tenant, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListTenants(ctx context.Context, arg ListTenantsParams) ([]Tenant, error)
}

var _ Querier = (*Queries)(nil)
