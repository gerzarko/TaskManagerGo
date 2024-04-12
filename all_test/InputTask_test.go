package all_test

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"taskManager/cmd/db"
)

type AllTaskToTest struct {
	AllTasks []db.Task
}

func TestInputTask(t *testing.T) {
	// _ = godotenv.Load()

	password := os.Getenv("all")
	dbNew := db.InitDatabase(password)
	defer db.CloseDatabase(dbNew)

	if password == "" {
		t.Errorf("this is an error of initializing db %v", password)
	}
	db.CreateTask(dbNew, "contenido", "2020/10/10", 3)
	taskReturned, _ := db.SelectTaskByContent(dbNew, "contenido")
	if taskReturned.Content != "contenido" {
		t.Errorf("error retreaving task from the db")
	}
}
