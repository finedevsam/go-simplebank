package db

import (
	"context"
	"testing"

	"github.com/samson/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        float64(util.RandomMoney()),
	}
	transfer, err := TestQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, account1.ID)
	require.Equal(t, arg.ToAccountID, account2.ID)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := TestQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
