package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lostak/amm/x/ocean/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// get pool id from tokens in message
	poolIndex, err := types.PoolIndex(msg.TokenA, msg.TokenB)
	if err != nil {
		return &types.MsgCreatePoolResponse{}, err
	}
	// if a pool is found, return an error
	_, found := k.GetPool(ctx, poolIndex)
	if found {
		return &types.MsgCreatePoolResponse{poolIndex}, errors.New("pool already exists")
	}
	// if token A is invalid return an error
	if !msg.TokenA.IsValid() {
		return &types.MsgCreatePoolResponse{}, errors.New("token a invalid")
	}
	// if token B is invalid return an error
	if !msg.TokenB.IsValid() {
		return &types.MsgCreatePoolResponse{}, errors.New("token B invalid")
	}

	pool := types.Pool{
		Index:		poolIndex,
		TokenA:		msg.TokenA,
		TokenB:		msg.TokenB,
		Shares:		msg.Shares,
		SwapFee:	msg.SwapFee,
	}

	// add pool to store from its index
	k.SetPool(ctx, pool)

	return &types.MsgCreatePoolResponse{}, nil
}
