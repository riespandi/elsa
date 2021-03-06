package cmd

import (
	"github.com/spf13/cobra"
  "github.com/matishsiao/goInfo"
)

func init() {
	rootCmd.AddCommand(vhostCreateCmd)
}

var vhostCreateCmd = &cobra.Command{
	Use:   "site",
	Short: "Generate and configure VirtualHost configuration file.",
	Run: func(cmd *cobra.Command, args []string) {
        gi := goInfo.GetInfo()
        gi.VarDump()
	},
}
