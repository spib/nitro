package cmd

import (
	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/internal/nitro"
)

var startCommand = &cobra.Command{
	Use:     "start",
	Aliases: []string{"up"},
	Short:   "Start a machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := flagMachineName

		startAction, err := nitro.Start(machine)
		if err != nil {
			return err
		}

		return nitro.Run(nitro.NewMultipassRunner("multipass"), []nitro.Action{*startAction})
	},
}