/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Siddheshk02/Securelee/auth"
	"github.com/Siddheshk02/Securelee/lib"
	"github.com/apoorvam/goterminal"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Securelee using your browser.",
	Long:  `Login to Securelee using your browser.`,
	Run: func(cmd *cobra.Command, args []string) {
		bl, re := auth.Check()
		c := color.New(color.FgYellow)

		if bl {
			c.Println("\n> " + re)
			c.Println("\n> Log-out to log-in/Sign-up with other Email.")
			c.Print("\n")
			os.Exit(1)
		}
		var email string

		c.Print("\n> Enter the Username (Email) : \n> ")
		fmt.Scanf("%s\n", &email)
		if len(strings.TrimSpace(email)) == 0 {
			err := fmt.Errorf("Your Email can't be empty %v", email)
			fmt.Println(err.Error())
			os.Exit(1)
		}
	reenter:
		c.Print("> Enter the Password (Minimum 8 characters) : \n> ")
		password, _ := term.ReadPassword(int(os.Stdin.Fd()))
		// fmt.Println("\033[8m") // Hide input
		// fmt.Scan(&password)
		fmt.Print("\n")
		if len(password) < 8 {
			fmt.Println("Enter a Valid Password!")
			goto reenter
		}
		fmt.Print("\n")

		code := lib.Mail(email)
		var check int
		c.Print("> To Verify the Email Address enter the code sent on the entered email address : ")
		fmt.Scanf("%d\n", &check)
		if check != code {
			c.Println("Invalid Code Entered!")
			os.Exit(1)
		}

		writer := goterminal.New(os.Stdout)
		ct.Foreground(ct.Yellow, false)
		for i := 0; i < 100; i = i + 10 {
			// add your text to writer's buffer
			fmt.Fprintf(writer, "> Authenticating (%d/100) ......\n", i)
			// write to terminal
			writer.Print()
			time.Sleep(time.Millisecond * 200)

			// clear the text written by the previous write, so that it can be re-written.
			writer.Clear()
		}
		writer.Reset()

		ct.Foreground(ct.Magenta, false)

		res := auth.Login(email, string(password))
		_ = res

		c = color.Set(color.FgCyan)
		c.Println("Verified!!")
		ct.ResetColor()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
