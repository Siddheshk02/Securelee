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

// SignUpCmd represents the SignUp command
var SignUpCmd = &cobra.Command{
	Use:   "SignUp",
	Short: "Sign-up to Securelee using your email.",
	Long:  `Sign-up to Securelee using your email.`,
	Run: func(cmd *cobra.Command, args []string) {

		bl, re := auth.Check()
		c := color.New(color.FgYellow)

		if bl {
			c.Println("\n> " + re)
			c.Println("\n> Log-out to log-in/Sign-up with other Email.")
			c.Print("\n")
			os.Exit(1)
		}

		var email, name string
		// var password string

		c.Print("\n> Enter the Email : \n> ")
		fmt.Scanf("%s\n", &email)
		if len(strings.TrimSpace(email)) == 0 {
			err := fmt.Errorf("Your Email can't be empty %v", email)
			fmt.Println(err.Error())
			os.Exit(1)
		}
	reenter:
		c.Print("> Enter the Password (Minimum 8 characters) : \n> ")
		password, _ := term.ReadPassword(int(os.Stdin.Fd()))

		fmt.Print("\n")
		if len(password) < 8 {
			fmt.Println("Enter a Valid Password!")
			goto reenter
		}

		c.Print("\n> Enter the Username : \n> ")
		fmt.Scanf("%s\n", &name)
		if len(strings.TrimSpace(name)) == 0 {
			err := fmt.Errorf("Your Username can't be empty %v", email)
			fmt.Println(err.Error())
			os.Exit(1)
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
			if i < 35 {
				fmt.Fprintf(writer, "> Authenticating (%d/100) ......\n", i)
			} else if i >= 35 && i < 70 {
				fmt.Fprintf(writer, "> Verifying email (%d/100) ......\n", i)
			} else {
				fmt.Fprintf(writer, "> Adding data (%d/100) ......\n", i)
			}

			writer.Print()
			time.Sleep(time.Millisecond * 200)

			writer.Clear()
		}
		writer.Reset()

		ct.Foreground(ct.Magenta, false)

		res := auth.Register(email, password, name)
		_ = res

		c = color.Set(color.FgCyan)
		c.Println("Verified!!")
		ct.ResetColor()
	},
}

func init() {
	rootCmd.AddCommand(SignUpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// SignUpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SignUpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
