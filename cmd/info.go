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

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, err := cmd.Flags().GetString("endpoint")
		if err != nil {
			fmt.Println("Error:", err)
		}
		if endpoint != "" {
			bk := bookieinfo(endpoint)
			var bkinfo BookieInfo
			json.Unmarshal(bk, &bkinfo)
			fmt.Printf("The FreeSpace is %v\n", bkinfo.FreeSpace)
			fmt.Printf("The TotalSpace is %v\n", bkinfo.TotalSpace)
		} else {
			fmt.Println("Plesase Input the bookie Endpoint address")
			os.Exit(1)
		}
	},
}

func init() {
	bookieCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
