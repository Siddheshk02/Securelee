/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Siddheshk02/Securelee/lib"
	"github.com/spf13/cobra"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share a File",
	Long:  `Upload a File -> Get a code -> Share the code with others.`,
	Run: func(cmd *cobra.Command, args []string) {
		var file string
		fmt.Print("\n> Enter the File-Path of the File to be Uploaded : ")
		fmt.Scanf("%s\n", &file)

		res := lib.Upload(file)

		fmt.Println(res)

	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
