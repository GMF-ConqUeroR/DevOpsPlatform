package main

import (
	"cmdb/cmd"

	"github.com/infraboard/mcube/validator"
	"github.com/spf13/cobra"
)

func main() {
	err := validator.Init()
	cobra.CheckErr(err)

	err = cmd.Execute()
	cobra.CheckErr(err)
}
