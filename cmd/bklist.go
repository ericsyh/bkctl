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
var bklistCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

bkctl is a personal demo CLI tool to learn Golang and Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, err := rootCmd.Flags().GetString("endpoint")
		if err != nil {
			fmt.Println("Error:", err)
		}
		if endpoint != "" {
			bk := bookielist(endpoint)
			json.Unmarshal(bk, &bklist)
			for i, _ := range bklist {
				fmt.Printf("Bookie Address is %s\n", i)
			}
		} else {
			fmt.Println("Plesase Input the bookie Endpoint address")
			os.Exit(1)
		}
	},
}

func init() {
	bookieCmd.AddCommand(bklistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
