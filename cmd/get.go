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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a File.",
	Long:  `Enter the Code -> Get the File.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Get called")
		bl, _ := auth.Check()
		c := color.New(color.FgYellow)

		if !bl {
			c.Println("\n> No User Logged-in.")
			c.Println("\n> Log-in or Sign-up to Securelee to continue.")
			c.Print("\n")
			os.Exit(1)
		}

		var code int
		c.Print("\n> Enter the code : ")
		fmt.Scanf("%d\n", &code)

		res := lib.Download(code)

		c.Print("\n> " + res + "\n\n")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
