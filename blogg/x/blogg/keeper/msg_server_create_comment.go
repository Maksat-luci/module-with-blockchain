package keeper

import (
	"context"
	"fmt"
	"strconv"

	"blogg/x/blogg/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("keys %d doesn't exist", msg.Id))
	}

	var comment = types.Comment{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Body:      msg.Body,
		Title:     msg.Title,
		PostID:    msg.PostID,
		CreatedAt: ctx.BlockHeight(),
	}
	temporaryVariable, err := strconv.Atoi(post.Createdat)
	if err != nil {
		return nil, fmt.Errorf("err: %s  %s", err.Error(),"not convertable")
	} 

	if comment.CreatedAt > int64(temporaryVariable)+100 {
		return nil, sdkerrors.Wrapf(types.ErrCommentOld, "Comment created at %d is older than post created at %d", comment.CreatedAt, post.Createdat)
	}

	id := k.AppendComment(ctx, comment)
	return &types.MsgCreateCommentResponse{Id: id}, nil
}
