package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mytestlab123/iochain3/testutil/keeper"
	"github.com/mytestlab123/iochain3/x/iochain3/keeper"
	"github.com/mytestlab123/iochain3/x/iochain3/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Iochain3Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
