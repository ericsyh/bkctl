/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ledgerlistCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, err := cmd.Flags().GetString("endpoint")
		if err != nil {
			fmt.Println("Error:", err)
		}
		if endpoint != "" {
			ledger := ledgerlist(endpoint)
			json.Unmarshal(ledger, &lglist)
			for i, _ := range lglist {
				fmt.Printf("Ledger Id is %s\n", i)
			}
		} else {
			fmt.Println("Plesase Input the bookie Endpoint address")
			os.Exit(1)
		}
	},
}

func init() {
	ledgerCmd.AddCommand(ledgerlistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
