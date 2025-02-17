package cmd

import (
	"auth/cmd/start"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "auth-api",
	Short: "用户中心",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// 命令具体的执行逻辑
		cmd.Help()
	},
}

func Execute() error {
	return rootCMD.Execute()
}

func init() {
	rootCMD.AddCommand(start.CMD)
}
