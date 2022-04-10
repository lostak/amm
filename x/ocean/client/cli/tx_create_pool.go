package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lostak/amm/x/ocean/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [token-a] [token-b] [shares] [swap-fee]",
		Short: "Broadcast message create-pool",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenA, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argTokenB, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			argShares, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			argSwapFee := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argShares,
				argSwapFee,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
