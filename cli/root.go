package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "clould-station",
	Version: "0.1",
	Short:   "云中转站",
	Long:    "云中转站，基于aliyun",
	Example: "clould-station -p ",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(1111)
		return nil
	},
}

func init() {
	// f := RootCmd.PersistentFlags()
	// f.StringVarP(&ossProvider, "provider", "p", "aliyun", "provider like aliyun...")
	// RootCmd.AddCommand()
}
