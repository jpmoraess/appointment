package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/jpmoraess/appointment-api/db/mock"
	db "github.com/jpmoraess/appointment-api/db/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	arg := db.RegisterTxParams{
		Name:     "john",
		Password: "123",
	}

	result := &db.RegisterTxResult{
		TenantID: 1,
		UserID:   2,
	}

	store.EXPECT().
		RegisterTx(gomock.Any(), gomock.Eq(arg)).
		Times(1).
		Return(result, nil)

	uc := NewRegister(store)

	err := uc.Execute(context.Background(), &RegisterInput{
		Name:     "john",
		Password: "123",
	})

	require.NoError(t, err)
	require.NotNil(t, result.TenantID)
	require.NotNil(t, result.UserID)
	require.Equal(t, int64(1), result.TenantID)
	require.Equal(t, int64(2), result.UserID)
}

func TestExecuteTableDriven(t *testing.T) {
	tests := []struct {
		name          string
		params        db.RegisterTxParams
		mockError     error
		mockReturn    *db.RegisterTxResult
		expectedError string
	}{
		{
			name: "Success",
			params: db.RegisterTxParams{
				Name:     "john",
				Password: "123",
			},
			mockError: nil,
			mockReturn: &db.RegisterTxResult{
				TenantID: 1,
				UserID:   2,
			},
			expectedError: "",
		},
		{
			name: "Error",
			params: db.RegisterTxParams{
				Name:     "john",
				Password: "123",
			},
			mockError:     errors.New("some error"),
			mockReturn:    nil,
			expectedError: "some error",
		},
	}

	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)
	defer ctrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store.EXPECT().
				RegisterTx(gomock.Any(), gomock.Eq(tt.params)).
				Times(1).
				Return(tt.mockReturn, tt.mockError)

			uc := NewRegister(store)

			err := uc.Execute(context.Background(), &RegisterInput{
				Name:     tt.params.Name,
				Password: tt.params.Password,
			})

			if len(tt.expectedError) == 0 {
				require.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}
