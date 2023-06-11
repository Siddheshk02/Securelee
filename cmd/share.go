/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Siddheshk02/Securelee/auth"
	"github.com/Siddheshk02/Securelee/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share a File",
	Long:  `Upload a File -> Get a code -> Share the code with others.`,
	Run: func(cmd *cobra.Command, args []string) {

		bl, _ := auth.Check()
		c := color.New(color.FgYellow)

		if !bl {
			c.Println("\n> No User Logged-in.")
			c.Println("\n> Log-in or Sign-up to Securelee to continue.")
			c.Print("\n")
			os.Exit(1)
		}

		var file string
		c.Print("\n> Enter the File-Path of the File to be Uploaded : ")
		fmt.Scanf("%s\n", &file)

		res := lib.Upload(file)

		c.Println("\n> " + res)
		fmt.Print("\n")

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
