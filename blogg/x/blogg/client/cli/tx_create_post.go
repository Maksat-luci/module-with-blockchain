package cli

import (
	"strconv"

	"blogg/x/blogg/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-post [title] [body] [createdat] [approved]",
		Short: "Broadcast message createPost",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argBody := args[1]
			argCreatedat := args[2]
			argApproved := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePost(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argBody,
				argCreatedat,
				argApproved,
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
