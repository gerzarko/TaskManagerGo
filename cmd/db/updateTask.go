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

var (
	idOfTask    int
	updContent  string
	priorityUpd int
)

func UpdateTaskPriority(db *sql.DB, idTask int, priorityUpdate int) int {
	task, err := SelectTaskById(db, idOfTask)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	fmt.Println("Old task ", task)

	if task.Id < 1 {
		fmt.Print("Task not found")
		return 0
	}

	// queryToSearch := "UPDATE test2.users SET first_name = ? WHERE (id_users = ?);"
	queryToSearch := "UPDATE `task`.`all_tasks` SET `priority` = ? WHERE (`id` = ?);"
	stmt, err := db.Prepare(queryToSearch)
	if err != nil {
		return 0
	}

	defer stmt.Close()

	qry, err := stmt.Exec(priorityUpdate, idOfTask)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	affectedR, err := qry.RowsAffected()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("RowsAffected ", affectedR)

	task, err = SelectTaskById(db, idOfTask)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	fmt.Println("upd task priority:", task.Priority)
	return 2
}

func UpdateTaskContent(db *sql.DB, idTask int, updContentTask string) int {
	task, err := SelectTaskById(db, idOfTask)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	fmt.Println("Old task ", task)

	if task.Id < 1 {
		fmt.Print("Task not found")
		return 0
	}

	// queryToSearch := "UPDATE test2.users SET first_name = ? WHERE (id_users = ?);"
	queryToSearch := "UPDATE `task`.`all_tasks` SET `content` = ? WHERE (`id` = ?);"
	stmt, err := db.Prepare(queryToSearch)
	if err != nil {
		return 0
	}

	defer stmt.Close()

	qry, err := stmt.Exec(updContentTask, idOfTask)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	affectedR, err := qry.RowsAffected()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("RowsAffected ", affectedR)

	task, err = SelectTaskById(db, idOfTask)
	if err != nil {
		fmt.Print(err)
		return 0
	}
	fmt.Println("upd task content:", task.Content)
	return 2
}

// updateTaskCmd represents the updateTask command
var updateTaskCmd = &cobra.Command{
	Use:   "updateTask",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateTask called")

		if idOfTask != 0 && updContent != "" {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			password := os.Getenv("all")
			db := InitDatabase(password)
			defer CloseDatabase(db)

			_ = UpdateTaskContent(db, idOfTask, updContent)
			return
		}

		if idOfTask != 0 && priorityUpd != 0 {

			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			password := os.Getenv("all")
			db := InitDatabase(password)
			defer CloseDatabase(db)

			_ = UpdateTaskPriority(db, idOfTask, priorityUpd)
			return
		}

		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	updateTaskCmd.Flags().IntVarP(&idOfTask, "idt", "u", 0, "the id to identify the task")
	updateTaskCmd.Flags().
		StringVarP(&updContent, "cont", "c", "", "the upd content of the task")

	updateTaskCmd.Flags().
		IntVarP(&priorityUpd, "prio", "p", 0, "the upd content of the task")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
