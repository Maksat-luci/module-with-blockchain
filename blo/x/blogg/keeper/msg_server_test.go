package keeper_test

import (
	"context"
	"testing"

	keepertest "blogg/testutil/keeper"
	"blogg/x/blogg/keeper"
	"blogg/x/blogg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BloggKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
