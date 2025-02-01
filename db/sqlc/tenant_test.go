package db

import (
	"context"
	"github.com/jpmoraess/appointment-api/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateTenant(t *testing.T) {
	var (
		name = util.RandomString(6)
		slug = util.RandomString(6)
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
