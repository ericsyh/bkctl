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

// diskfileCmd represents the diskfile command
var diskfileCmd = &cobra.Command{
	Use:   "diskfile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, err := cmd.Flags().GetString("endpoint")
		if err != nil {
			fmt.Println("Error:", err)
		}
		if endpoint != "" {
			bk := diskfile(endpoint)
			var bkdisk DiskFile
			json.Unmarshal(bk, &bkdisk)
			fmt.Printf("The IndexFiles is %v\n", bkdisk.IndexFiles)
			fmt.Printf("The JournalFiles is %v\n", bkdisk.JournalFiles)
			fmt.Printf("The EntrylogFiles is %v\n", bkdisk.EntrylogFiles)
		} else {
			fmt.Println("Plesase Input the bookie Endpoint address")
			os.Exit(1)
		}
	},
}

func init() {
	bookieCmd.AddCommand(diskfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
