package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group mbgovperm queries under a subcommand
	mbgovpermQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	mbgovpermQueryCmd.AddCommand(
		flags.GetCommands(
	// this line is used by starport scaffolding # 1
			GetCmdListConsent(queryRoute, cdc),
			GetCmdGetConsent(queryRoute, cdc),
			GetCmdListExtservice(queryRoute, cdc),
			GetCmdGetExtservice(queryRoute, cdc),
	// TODO: Add query Cmds
		)...,
	)

	return mbgovpermQueryCmd
}

// TODO: Add Query Commands
