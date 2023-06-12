/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Siddheshk02/Securelee/auth"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Command to Logout of Securelee.",
	Long:  `Command to Logout of Securelee.`,
	Run: func(cmd *cobra.Command, args []string) {

		bl, _ := auth.Check()
		c := color.New(color.FgYellow)

		if !bl {
			c.Println("\n> No User Logged-in.")
			c.Println("\n> Log-in or Sign-up to Securelee to continue.")
			c.Print("\n")
			os.Exit(1)
		}

		res := auth.Logout()

		c.Println("\n> " + res + "\n")

	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
