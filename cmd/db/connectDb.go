/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userPassword string

// connectDbCmd represents the connectDb command
var ConnectDbCmd = &cobra.Command{
	Use:   "connectDb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if userPassword == "" {
			err := cmd.Help()
			if err != nil {
				fmt.Print(err)
			}
		}
		fmt.Print(userPassword)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	ConnectDbCmd.Flags().StringVarP(&userPassword, "login", "l", "", "login db ")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectDbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectDbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
