/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// metaCmd represents the meta command
var (
	ledgerid int
	metaCmd  = &cobra.Command{
		Use:   "meta",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetInt("ledgerid")
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(id)
		},
	}
)

func init() {
	ledgerCmd.AddCommand(metaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	metaCmd.Flags().IntVarP(&ledgerid, "ledgerid", "l", 0, "Ledger Id to get Metadata")
	metaCmd.MarkFlagRequired("ledgerid")
}
