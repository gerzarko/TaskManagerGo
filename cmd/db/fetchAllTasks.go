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
	isOlderFirsts      bool
	lowerPriorityFirst bool
)

func SelectAlltasks(db *sql.DB) (tasksFound []Task, err error) {
	queryToDo := "SELECT * FROM task.all_tasks;"

	stament, erro := db.Prepare(queryToDo)

	if erro != nil {
		log.Fatal(erro)
	}
	defer stament.Close()
	// row, err := db.Query(queryToDo)
	row, err := stament.Query()
	if err != nil {
		fmt.Print("error searching for the database")
		return tasksFound, err
	}

	// next devuelve true mientras haya otra row valida
	for row.Next() {
		taskFound := Task{}
		err := row.Scan(
			&taskFound.Id,
			&taskFound.Content,
			&taskFound.Date,
			&taskFound.Priority,
		)
		if err != nil {
			return tasksFound, err
		}
		tasksFound = append(tasksFound, taskFound)
	}
	return tasksFound, nil
}

func SelectAllTasksOldFirst(db *sql.DB) (tasksFound []Task, err error) {
	queryToDo := "SELECT * FROM task.all_tasks ORDER BY date ASC;"

	stament, erro := db.Prepare(queryToDo)

	if erro != nil {
		log.Fatal(erro)
	}
	defer stament.Close()
	// row, err := db.Query(queryToDo)
	row, err := stament.Query()
	if err != nil {
		fmt.Print("error searching for the database")
		return tasksFound, err
	}

	// next devuelve true mientras haya otra row valida
	for row.Next() {
		taskFound := Task{}
		err := row.Scan(
			&taskFound.Id,
			&taskFound.Content,
			&taskFound.Date,
			&taskFound.Priority,
		)
		if err != nil {
			return tasksFound, err
		}
		tasksFound = append(tasksFound, taskFound)
	}
	return tasksFound, nil
}

func SelectAllTasksLowPriorityFirst(db *sql.DB) (tasksFound []Task, err error) {
	queryToDo := "SELECT * FROM task.all_tasks ORDER BY priority ASC;"

	stament, erro := db.Prepare(queryToDo)

	if erro != nil {
		log.Fatal(erro)
	}
	defer stament.Close()
	// row, err := db.Query(queryToDo)
	row, err := stament.Query()
	if err != nil {
		fmt.Print("error searching for the database")
		return tasksFound, err
	}

	// next devuelve true mientras haya otra row valida
	for row.Next() {
		taskFound := Task{}
		err := row.Scan(
			&taskFound.Id,
			&taskFound.Content,
			&taskFound.Date,
			&taskFound.Priority,
		)
		if err != nil {
			return tasksFound, err
		}
		tasksFound = append(tasksFound, taskFound)
	}
	return tasksFound, nil
}

// fetchAllTasksCmd represents the fetchAllTasks command
var FetchAllTasksCmd = &cobra.Command{
	Use:   "fetchAllTasks",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("fetchAllTasks called")
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		password := os.Getenv("all")
		db := InitDatabase(password)

		defer CloseDatabase(db)

		if isOlderFirsts {
			tasks, err := SelectAllTasksOldFirst(db)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(isOlderFirsts)
			fmt.Println(tasks)
			return
		}
		if lowerPriorityFirst {
			tasks, err := SelectAllTasksLowPriorityFirst(db)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(isOlderFirsts)
			fmt.Println(tasks)
			return
		} else {
			tasks, err := SelectAlltasks(db)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(tasks)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	FetchAllTasksCmd.Flags().
		BoolVarP(&isOlderFirsts, "old", "o", false, "for older tasks first")

	FetchAllTasksCmd.Flags().
		BoolVarP(&lowerPriorityFirst, "low", "l", false, "for lower priority tasks first")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchAllTasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchAllTasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
