/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type Task struct {
	Id       int
	Content  string
	Date     string
	Priority int
}

func InitDatabase(userPassword string) *sql.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// password := os.Getenv("all")

	db, err := sql.Open("mysql", userPassword)
	if err != nil {
		fmt.Print("error entering the database\n")
		fmt.Print(err)
	}

	// defer db.Close()
	// check if database connection to db is fine
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CloseDatabase(db *sql.DB) {
	db.Close()
}

// dbCmd represents the db command
var DbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database Commands",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("db called")
	},
}

func init() {
	DbCmd.AddCommand(ConnectDbCmd)
	DbCmd.AddCommand(FetchTaskCmd)
	DbCmd.AddCommand(FetchAllTasksCmd)
	DbCmd.AddCommand(updateTaskCmd)
	DbCmd.AddCommand(deleteTaskCmd)
	DbCmd.AddCommand(createTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
