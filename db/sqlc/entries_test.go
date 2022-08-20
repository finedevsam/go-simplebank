package db

import (
	"context"
	"testing"

	"github.com/samson/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntries(t *testing.T) Entry {
	account1 := CreateRandomAccount(t)
	arg := CreateEntriesParams{
		AccountID: account1.ID,
		Amount:    float64(util.RandomMoney()),
	}
	entry, err := TestQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, account1.ID)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry

}
func TestCreateEntries(t *testing.T) {
	CreateRandomEntries(t)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntries(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := TestQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
