package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/peggy/x/ethbridge/types"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdCreateEthBridgeClaim is the CLI command for creating a claim on an ethereum prophecy
func GetCmdCreateEthBridgeClaim(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-claim [chain-id] [nonce] [ethereum-sender-address] [cosmos-receiver-address] [validator-address] [amount]",
		Short: "create a claim on an ethereum prophecy",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			chainID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			nonce, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			ethereumSender := types.NewEthereumAddress(args[2])

			cosmosReceiver, err := sdk.AccAddressFromBech32(args[3])
			if err != nil {
				return err
			}

			validator, err := sdk.ValAddressFromBech32(args[4])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoins(args[5])
			if err != nil {
				return err
			}

			ethBridgeClaim := types.NewEthBridgeClaim(chainID, ethereumSender, nonce, "eth", ethereumSender, ethereumSender, cosmosReceiver, validator, amount)
			msg := types.NewMsgCreateEthBridgeClaim(ethBridgeClaim)

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}
