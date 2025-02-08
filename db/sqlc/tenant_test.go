package db

import (
	"context"
	"testing"

	"github.com/jpmoraess/appointment-api/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestQueries_CreateTenant(t *testing.T) {
	var (
		name = utils.RandomString(6)
		slug = utils.RandomString(6)
	)

	arg := CreateTenantParams{
		Name: name,
		Slug: slug,
	}

	tenant, err := testQueries.CreateTenant(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tenant)
	require.Equal(t, name, tenant.Name)
	require.Equal(t, slug, tenant.Slug)
}
