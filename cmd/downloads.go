/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Siddheshk02/Securelee/auth"
	"github.com/Siddheshk02/Securelee/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// downloadsCmd represents the downloads command
var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "User who downloaded the Files shared by you",
	Long:  `User who downloaded the Files shared by you.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("downloads called")
		bl, _ := auth.Check()
		c := color.New(color.FgYellow)

		if !bl {
			c.Println("\n> No User Logged-in.")
			c.Println("\n> Log-in or Sign-up to Securelee to continue.")
			c.Print("\n")
			os.Exit(1)
		}

		lib.Getdownloads()
	},
}

func init() {
	rootCmd.AddCommand(downloadsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
