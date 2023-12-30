package cli

import (
	"strconv"

	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPostResult() *cobra.Command {
	//cmd := &cobra.Command{
	//	Use:   "post-result [result]",
	//	Short: "Broadcast message post-result",
	//	Args:  cobra.ExactArgs(1),
	//	RunE: func(cmd *cobra.Command, args []string) (err error) {
	//		clientCtx, err := client.GetClientTxContext(cmd)
	//		if err != nil {
	//			return err
	//		}
	//
	//		num:=strconv.Atoi(args[0]),
	//		if num!=nil{
	//			ret
	//		}}
	//		msg := types.NewMsgPostResult(
	//			clientCtx.GetFromAddress().String(),
	//
	//		strconv.Atoi(args[1)
	//		)
	//		if err := msg.ValidateBasic(); err != nil {
	//			return err
	//		}
	//		return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
	//	},
	//}

	//flags.AddTxFlagsToCmd(cmd)

	//return cmd
	return nil
}
