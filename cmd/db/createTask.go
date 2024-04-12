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
	newContent  string
	newDate     string
	newPriority int
)

// Creates a new task and push it to the DB
func CreateTask(db *sql.DB, content string, date string, priority int) {
	queryTodo := "INSERT INTO `task`.`all_tasks` (`id`, `content`, `date`, `priority`) VALUES (?, ?, ? ,?);"
	allTasks, err := SelectAlltasks(db)
	if err != nil {
		log.Fatal(err)
	}
	length := len(allTasks)
	idNewTask := length + 1
	// idNewTask := allTasks[len(allTasks)] + 1
	stmt, err := db.Prepare(queryTodo)
	if err != nil {
		fmt.Println(err)
	}

	qry, err := stmt.Exec(idNewTask, content, date, priority)
	if err != nil {
		fmt.Println(err)
	}

	raffec, err := qry.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Created and row affected ", raffec)
	newTask, err := SelectTaskById(db, idNewTask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newTask)
}

// createTaskCmd represents the createTask command
var createTaskCmd = &cobra.Command{
	Use:   "createTask",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createTask called")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		password := os.Getenv("all")
		db := InitDatabase(password)
		defer CloseDatabase(db)

		if newContent != "" {
			CreateTask(db, newContent, newDate, newPriority)
			return
		}
		err = cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	createTaskCmd.Flags().StringVarP(&newDate, "dat", "d", "", "the data for the new task")
	createTaskCmd.Flags().
		StringVarP(&newContent, "cont", "c", "", "the content for the new task")
	createTaskCmd.Flags().
		IntVarP(&newPriority, "prio", "p", 0, "the new content of the task")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
