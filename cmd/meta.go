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
			endpoint, err := cmd.Flags().GetString("endpoint")
			if err != nil {
				fmt.Println("Error:", err)
			}
			if endpoint != "" {
				ledger := ledgermeta(endpoint, id)
				var lgmeta LedgerMeta
				json.Unmarshal(ledger, &lgmeta)
				fmt.Printf("LedgerID is %v\n", lgmeta.ID.LedgerID)
				fmt.Printf("MetadataFormatVersion  is %v\n", lgmeta.ID.MetadataFormatVersion)
				fmt.Printf("EnsembleSize is %v\n", lgmeta.ID.EnsembleSize)
				fmt.Printf("WriteQuorumSize is %v\n", lgmeta.ID.WriteQuorumSize)
				fmt.Printf("AckQuorumSize is %v\n", lgmeta.ID.AckQuorumSize)
				fmt.Printf("State is %v\n", lgmeta.ID.State)
				fmt.Printf("Length is %v\n", lgmeta.ID.Length)
				fmt.Printf("LastEntryID is %v\n", lgmeta.ID.LastEntryID)
				fmt.Printf("Create Time is %v\n", lgmeta.ID.Ctime)
				for i := range lgmeta.ID.AllEnsembles.EnsembleId {
					fmt.Printf("EnsembleId is %s\n", lgmeta.ID.AllEnsembles.EnsembleId[i])
				}
			} else {
				fmt.Println("Plesase Input the bookie Endpoint address")
				os.Exit(1)
			}
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
