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

var idToSearch int

func SelectTaskById(db *sql.DB, id int) (taskFound Task, err error) {
	queryToDo := "SELECT * FROM task.all_tasks where id = ?;"

	// queryToDo := "SELECT * FROM test2.users WHERE first_name= ? AND last_name= ?;"
	stmt, erro := db.Prepare(queryToDo)

	if erro != nil {
		log.Fatal(erro)
	}
	defer stmt.Close()

	// Search in database the query and replace the interrogation sign for nameToSearch
	// It returns nothing or success, it doesn't return error
	// row := db.QueryRow(queryToDo, nameToSearch)
	row := stmt.QueryRow(id)

	// fmt.Print(row)

	// Assigns to the return var the data that it found
	err = row.Scan(&taskFound.Id, &taskFound.Content, &taskFound.Date, &taskFound.Priority)
	if err != nil {
		return taskFound, err
	}
	return taskFound, nil
}

func SelectTaskByContent(db *sql.DB, contenido string) (taskFound Task, err error) {
	queryToDo := "SELECT * FROM task.all_tasks where content = ?;"

	// queryToDo := "SELECT * FROM test2.users WHERE first_name= ? AND last_name= ?;"
	stmt, erro := db.Prepare(queryToDo)

	if erro != nil {
		log.Fatal(erro)
	}
	defer stmt.Close()

	// Search in database the query and replace the interrogation sign for nameToSearch
	// It returns nothing or success, it doesn't return error
	// row := db.QueryRow(queryToDo, nameToSearch)
	row := stmt.QueryRow(contenido)

	// fmt.Print(row)

	// Assigns to the return var the data that it found
	err = row.Scan(&taskFound.Id, &taskFound.Content, &taskFound.Date, &taskFound.Priority)
	if err != nil {
		return taskFound, err
	}
	return taskFound, nil
}

// fetchTaskCmd represents the fetchTask command
var FetchTaskCmd = &cobra.Command{
	Use:   "fetchTask",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetchTask called")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		password := os.Getenv("all")
		db := InitDatabase(password)

		defer CloseDatabase(db)
		fmt.Println(idToSearch)

		task, err := SelectTaskById(db, idToSearch)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(task)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	FetchTaskCmd.Flags().IntVarP(&idToSearch, "id", "i", 0, "the id to search in the task list")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
