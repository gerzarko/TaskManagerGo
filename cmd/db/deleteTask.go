/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var idOfTaskToDel int

func DeleteTask(db *sql.DB, idOfTask int) int {
	queryToDo := "DELETE FROM task.all_tasks WHERE (id= ?)"
	stmt, err := db.Prepare(queryToDo)
	if err != nil {
		fmt.Print("cant do the prepare\n")
		return 0
	}

	task, err := SelectTaskById(db, idOfTask)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	qry, err := stmt.Exec(task.Id)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	affectedR, err := qry.RowsAffected()
	if err != nil {
		fmt.Print(err)
		return 0
	}
	fmt.Println("RowsAffected: ", affectedR)
	return 1
}

// deleteTaskCmd represents the deleteTask command
var deleteTaskCmd = &cobra.Command{
	Use:   "deleteTask",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteTask called")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		password := os.Getenv("all")
		db := InitDatabase(password)

		defer CloseDatabase(db)
		if idOfTaskToDel != 0 && idOfTaskToDel > 0 {
			DeleteTask(db, idOfTaskToDel)
			return
		}
		err = cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	deleteTaskCmd.Flags().IntVarP(&idOfTaskToDel, "del", "d", 0, "id of task to delete")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
