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

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Get list of Files Shared by you.",
	Long:  `Get list of Files Shared by you.`,
	Run: func(cmd *cobra.Command, args []string) {
		bl, _ := auth.Check()
		c := color.New(color.FgYellow)

		if !bl {
			c.Println("\n> No User Logged-in.")
			c.Println("\n> Log-in or Sign-up to Securelee to continue.")
			c.Print("\n")
			os.Exit(1)
		}

		lib.View()

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
