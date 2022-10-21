package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochain3/testutil/keeper"
	"github.com/mytestlab123/iochain3/x/iochain3/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Iochain3Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
