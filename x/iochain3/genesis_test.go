package iochain3_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochain3/testutil/keeper"
	"github.com/mytestlab123/iochain3/testutil/nullify"
	"github.com/mytestlab123/iochain3/x/iochain3"
	"github.com/mytestlab123/iochain3/x/iochain3/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Iochain3Keeper(t)
	iochain3.InitGenesis(ctx, *k, genesisState)
	got := iochain3.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
