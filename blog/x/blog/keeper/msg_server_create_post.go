package keeper

import (
	"context"
	"time"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message

	var post = types.Post{
		Creator:   msg.Creator,
		Title:     msg.Title,
		Body:      msg.Body,
		Createdat: time.Now().Format("2006-01-02 15:04:05"),
		Approved:  true,
	}
	id := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{Id: id}, nil
}
