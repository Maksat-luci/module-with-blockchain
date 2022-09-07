package keeper

import (
	"context"

	"blogg/x/blogg/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var post = types.Post{
		Creator:   msg.Creator,
		Title:     msg.Title,
		Body:      msg.Body,
		Createdat: ctx.BlockHeader().Time.Format("2006-01-02 15:04:05"),
		Approved:  true,
		
	}

	iD := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{Id: iD}, nil
}
